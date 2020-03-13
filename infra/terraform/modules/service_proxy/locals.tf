locals {
  ecs_family         = format("proxy-%s", var.TZ_NETWORK)
  ecs_service        = format("proxy-%s", var.TZ_NETWORK)
  launch_type        = "FARGATE"
  proxy_docker_image = "${var.PROXY_DOCKER_IMAGE_NAME}:${var.PROXY_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = format("tzlink-proxy-%s", var.TZ_NETWORK)

  network_certificate_arn = "arn:aws:acm:us-east-1:609827314188:certificate/0e984566-3c5d-4c41-b211-461bb4c1b845"
}