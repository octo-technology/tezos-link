resource "aws_cloudwatch_log_group" "proxy" {
  name              = format("tzlink-proxy-%s", var.TZ_NETWORK)
  retention_in_days = 7
}