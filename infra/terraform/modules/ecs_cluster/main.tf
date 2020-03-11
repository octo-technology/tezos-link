resource "aws_ecs_cluster" "cluster" {
  name = "tzlink"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}