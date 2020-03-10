resource "aws_cloudwatch_log_group" "proxy" {
  name              = "tzlink-proxy"
  retention_in_days = 7
}