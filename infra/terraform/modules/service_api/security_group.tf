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
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_lb.id
}

resource "aws_security_group" "api_ecs_task" {
  name        = "api_ecs_task"
  description = "Security group for api (access only by loadbalancer)"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "main_ingress_for_api" {

  type                     = "ingress"
  from_port                = var.API_PORT
  to_port                  = var.API_PORT
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.api_lb.id

  security_group_id = aws_security_group.api_ecs_task.id
}

resource "aws_security_group_rule" "all_egress_for_api" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.api_ecs_task.id
}