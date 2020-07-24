terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

resource "aws_ecs_cluster" "cluster" {
  name = "tzlink"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}