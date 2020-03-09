resource "aws_ecs_cluster" "cluster" {
  name = format("tzlink-%s", var.ENV)
}