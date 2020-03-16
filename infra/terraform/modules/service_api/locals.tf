locals {
  ecs_family       = "api"
  ecs_service      = "api"
  launch_type      = "FARGATE"
  api_docker_image = "${var.API_DOCKER_IMAGE_NAME}:${var.API_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = "tzlink-api"
}