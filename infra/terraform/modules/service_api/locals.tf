locals {
  ecs_family       = "api"
  ecs_service      = "api"
  launch_type      = "FARGATE"
  api_docker_image = "${var.API_DOCKER_IMAGE_NAME}:${var.API_DOCKER_IMAGE_VERSION}"

  ecs_task_logs_stream_prefix = "tzlink-api"

  api_certificate_arn = "	arn:aws:acm:us-east-1:609827314188:certificate/d9adfdca-f0d7-43df-a560-088d6c98ba05"
}