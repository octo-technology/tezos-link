data "aws_iam_role" "tzlink_ecs_tasks_access" {
  name = "tzlink_ecs_tasks_access"
}

data "aws_db_instance" "database" {
  db_instance_identifier = format("tzlink-%s-database", var.ENV)
}

resource "aws_ecs_task_definition" "backend" {
  family                   = local.ecs_family
  network_mode             = "awsvpc"
  requires_compatibilities = [local.launch_type]
  cpu                      = var.BACKEND_CPU
  memory                   = var.BACKEND_MEMORY

  execution_role_arn = data.aws_iam_role.tzlink_ecs_tasks_access.arn

  container_definitions = templatefile("${path.module}/proxy_task_definition.json",
    {
      task_name   = local.ecs_family,
      task_image  = local.backend_docker_image,
      task_port   = var.BACKEND_PORT,
      task_cpu    = var.BACKEND_CPU,
      task_memory = var.BACKEND_MEMORY,

      database_url      = data.aws_db_instance.database.endpoint,
      database_username = var.DATABASE_USERNAME,
      database_password = var.DATABASE_PASSWORD,
      database_table    = data.aws_db_instance.database.db_name,

      environment_config = var.BACKEND_CONFIGURATION_FILE,

      log_group_name          = aws_cloudwatch_log_group.backend.name,
      log_group_region        = var.REGION,
      log_group_stream_prefix = local.ecs_task_logs_stream_prefix
  })
}

resource "aws_ecs_service" "backend" {
  name            = local.ecs_service
  cluster         = data.aws_ecs_cluster.backend.arn
  task_definition = aws_ecs_task_definition.backend.arn
  desired_count   = var.BACKEND_DESIRED_COUNT
  launch_type     = local.launch_type

  network_configuration {
    security_groups = [aws_security_group.backend_ecs_task.id]
    subnets         = tolist(data.aws_subnet_ids.tzlink_private_backend.ids)
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.backend.id
    container_name   = local.ecs_family
    container_port   = var.BACKEND_PORT
  }

  depends_on = [aws_alb_listener.backend]
}
