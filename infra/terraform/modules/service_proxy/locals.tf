locals {
  ecs_family         = format("proxy-%s", var.TZ_NETWORK)
  ecs_service        = format("proxy-%s", var.TZ_NETWORK)
  launch_type        = "FARGATE"
  proxy_docker_image = "${var.PROXY_DOCKER_IMAGE_NAME}:${var.PROXY_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = format("tzlink-proxy-%s", var.TZ_NETWORK)
}