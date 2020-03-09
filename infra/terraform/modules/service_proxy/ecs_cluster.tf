data "aws_ecs_cluster" "proxy" {
  cluster_name = format("tzlink-%s", var.ENV)
}