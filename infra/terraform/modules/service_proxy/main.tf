data "aws_iam_role" "tzlink_ecs_tasks_access" {
  name = "tzlink_ecs_tasks_access"
}

resource "aws_ecs_task_definition" "proxy" {
  family                   = local.ecs_family
  network_mode             = "awsvpc"
  requires_compatibilities = [local.launch_type]
  cpu                      = var.PROXY_CPU
  memory                   = var.PROXY_MEMORY

  execution_role_arn = data.aws_iam_role.tzlink_ecs_tasks_access.arn

  container_definitions = templatefile("${path.module}/proxy_task_definition.json",
    {
      task_name   = local.ecs_family,
      task_image  = local.proxy_docker_image,
      task_port   = var.PROXY_PORT,
      task_cpu    = var.PROXY_CPU,
      task_memory = var.PROXY_MEMORY,

      database_url      = var.DATABASE_URL
      database_username = var.DATABASE_USERNAME
      database_password = var.DATABASE_PASSWORD
      database_table    = var.DATABASE_TABLE

      tezos_host = var.TEZOS_FARM_URL
      tezos_port = var.TEZOS_FARM_PORT

      environment_config = var.PROXY_CONFIGURATION_FILE,

      log_group_name          = aws_cloudwatch_log_group.proxy.name,
      log_group_region        = var.REGION,
      log_group_stream_prefix = local.ecs_task_logs_stream_prefix
  })
}

resource "aws_ecs_service" "proxy" {
  name            = local.ecs_service
  cluster         = data.aws_ecs_cluster.proxy.arn
  task_definition = aws_ecs_task_definition.proxy.arn
  desired_count   = var.PROXY_DESIRED_COUNT
  launch_type     = local.launch_type

  network_configuration {
    security_groups = [aws_security_group.proxy_ecs_task.id]
    subnets         = tolist(data.aws_subnet_ids.tzlink_private_proxy.ids)
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.proxy.id
    container_name   = local.ecs_family
    container_port   = var.PROXY_PORT
  }

  depends_on = [aws_alb_listener.proxy]
}
