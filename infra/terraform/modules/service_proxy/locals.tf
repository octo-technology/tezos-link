locals {
  ecs_family         = "proxy"
  ecs_service        = "proxy"
  launch_type        = "FARGATE"
  proxy_docker_image = "${var.PROXY_DOCKER_IMAGE_NAME}:${var.PROXY_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = "tzlink-proxy"
}