resource "aws_cloudwatch_log_group" "backend" {
  name              = format("tzlink-backend-%s", var.ENV)
  retention_in_days = 7
}