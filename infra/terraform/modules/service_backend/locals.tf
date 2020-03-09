locals {
  ecs_family           = "backend"
  ecs_service          = "backend"
  launch_type          = "FARGATE"
  backend_docker_image = "${var.BACKEND_DOCKER_IMAGE_NAME}:${var.BACKEND_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = "tzlink-backend"
}