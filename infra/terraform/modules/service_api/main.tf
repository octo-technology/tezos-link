data "aws_iam_role" "tzlink_ecs_tasks_access" {
  name = "tzlink_ecs_tasks_access"
}

data "aws_db_instance" "database" {
  db_instance_identifier = "tzlink-database"
}

resource "aws_ecs_task_definition" "api" {
  family                   = local.ecs_family
  network_mode             = "awsvpc"
  requires_compatibilities = [local.launch_type]
  cpu                      = var.API_CPU
  memory                   = var.API_MEMORY

  execution_role_arn = data.aws_iam_role.tzlink_ecs_tasks_access.arn

  container_definitions = templatefile("${path.module}/proxy_task_definition.json",
    {
      task_name   = local.ecs_family,
      task_image  = local.api_docker_image,
      task_port   = var.API_PORT,
      task_cpu    = var.API_CPU,
      task_memory = var.API_MEMORY,

      database_url      = data.aws_db_instance.database.endpoint,
      database_username = var.DATABASE_USERNAME,
      database_password = var.DATABASE_PASSWORD,
      database_table    = data.aws_db_instance.database.db_name,

      environment_config = var.API_CONFIGURATION_FILE,

      log_group_name          = aws_cloudwatch_log_group.api.name,
      log_group_region        = var.REGION,
      log_group_stream_prefix = local.ecs_task_logs_stream_prefix
  })
}

resource "aws_ecs_service" "api" {
  name            = local.ecs_service
  cluster         = data.aws_ecs_cluster.cluster.arn
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = var.API_DESIRED_COUNT
  launch_type     = local.launch_type

  network_configuration {
    security_groups = [aws_security_group.api_ecs_task.id]
    subnets         = tolist(data.aws_subnet_ids.tzlink_private_ecs.ids)
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.api.id
    container_name   = local.ecs_family
    container_port   = var.API_PORT
  }

  depends_on = [aws_alb_listener.api]
}
