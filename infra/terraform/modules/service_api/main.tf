terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

# Manually created role based on iam/ecs-task-assume-role
data "aws_iam_role" "tzlink_ecs_tasks_access" {
  name = "tzlink_ecs_tasks_executor"
}

# Network configuration

data "aws_vpc" "tzlink" {
  cidr_block = var.vpc_cidr

  tags = {
    Name    = "tzlink"
    Project = var.project_name
  }
}

data "aws_subnet_ids" "tzlink_public_ecs" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-public-ecs-*"
    Project = var.project_name
  }
}

data "aws_subnet_ids" "tzlink_private_ecs" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-private-ecs-*"
    Project = var.project_name
  }
}

# Network access control

resource "aws_security_group" "api_lb" {
  name        = "api_loadbalancer"
  description = "Security group for loadbalancer"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_api_loadbalancer" {
  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_lb.id
}

resource "aws_security_group_rule" "https_ingress_for_api_loadbalancer" {
  type        = "ingress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_lb.id
}

resource "aws_security_group_rule" "all_egress_for_api_loadbalancer" {
  type                     = "egress"
  from_port                = var.port
  to_port                  = var.port
  protocol                 = "-1"
  source_security_group_id = aws_security_group.api_ecs_task.id

  security_group_id = aws_security_group.api_lb.id
}

resource "aws_security_group" "api_ecs_task" {
  name        = "api_ecs_task"
  description = "Security group for api (access only by loadbalancer)"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "main_ingress_for_api" {

  type                     = "ingress"
  from_port                = var.port
  to_port                  = var.port
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.api_lb.id

  security_group_id = aws_security_group.api_ecs_task.id
}

resource "aws_security_group_rule" "all_egress_for_api" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.api_ecs_task.id
}

resource "aws_security_group_rule" "http_egress_for_proxy" {
  type        = "egress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_ecs_task.id
}

resource "aws_security_group_rule" "https_egress_for_proxy" {
  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_ecs_task.id
}

# Public DNS record configuration

data "aws_route53_zone" "tezoslink" {
  name = "tezoslink.io."
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.tezoslink.zone_id
  name    = "api.tezoslink.io"
  type    = "A"

  alias {
    name                   = aws_alb.api.dns_name
    zone_id                = aws_alb.api.zone_id
    evaluate_target_health = false
  }
}

# RDS database

data "aws_rds_cluster" "database" {
  cluster_identifier = "tzlink-database"
}

# ECS service configuration

data "aws_ecs_cluster" "cluster" {
  cluster_name = "tzlink"
}

data "aws_secretsmanager_secret" "tzlink_database_password" {
  name = "tzlink-database-password"
}

resource "aws_ecs_task_definition" "api" {
  family                   = "api"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.cpu
  memory                   = var.memory

  execution_role_arn = data.aws_iam_role.tzlink_ecs_tasks_access.arn

  container_definitions = templatefile("${path.module}/templates/proxy_task_definition.json.tpl",
    {
      task_name  = "api",
      task_image = format("%s:%s", var.docker_image_name, var.docker_image_version),

      task_port   = var.port,
      task_cpu    = var.cpu,
      task_memory = var.memory,

      database_url          = data.aws_rds_cluster.database.endpoint,
      database_username     = var.database_master_username,
      database_password_arn = data.aws_secretsmanager_secret.tzlink_database_password.arn,
      database_name         = data.aws_rds_cluster.database.database_name,

      configuration_file = var.configuration_file,

      log_group_name          = aws_cloudwatch_log_group.api.name,
      log_group_region        = var.region,
      log_group_stream_prefix = "api"
  })
}

resource "aws_ecs_service" "api" {
  name            = "api"
  cluster         = data.aws_ecs_cluster.cluster.arn
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = var.desired_container_number
  launch_type     = "FARGATE"

  network_configuration {
    security_groups = [aws_security_group.api_ecs_task.id]
    subnets         = tolist(data.aws_subnet_ids.tzlink_private_ecs.ids)
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.api.id
    container_name   = "api"
    container_port   = var.port
  }

  depends_on = [aws_alb_listener.api_https]
}

# Loadbalancer configuration

resource "aws_alb" "api" {
  name            = "tzlink-api"
  subnets         = tolist(data.aws_subnet_ids.tzlink_public_ecs.ids)
  security_groups = [aws_security_group.api_lb.id]

  tags = {
    Name    = "tzlink-api"
    Project = var.project_name
  }
}

resource "aws_alb_target_group" "api" {
  name        = "tzlink-api"
  port        = var.port
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.tzlink.id
  target_type = "ip"

  stickiness {
    enabled         = true
    type            = "lb_cookie"
    cookie_duration = 600
  }

  health_check {
    enabled  = true
    path     = "/health"
    port     = var.port
    protocol = "HTTP"
  }

  tags = {
    Name    = "tzlink-api"
    Project = var.project_name
  }

  depends_on = [aws_alb.api]
}

data "aws_acm_certificate" "api" {
  domain   = "api.tezoslink.io"
  statuses = ["ISSUED"]
}

resource "aws_alb_listener" "api_https" {
  load_balancer_arn = aws_alb.api.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS-1-2-Ext-2018-06"
  certificate_arn   = data.aws_acm_certificate.api.arn

  default_action {
    target_group_arn = aws_alb_target_group.api.arn
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.api]
}

resource "aws_alb_listener" "api_http_redirect" {
  load_balancer_arn = aws_alb.api.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = 443
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }

  depends_on = [aws_alb_target_group.api]
}

# Cloudwatch Logs

resource "aws_cloudwatch_log_group" "api" {
  name              = "/aws/ecs/service/tzlink/api"
  retention_in_days = 7
}