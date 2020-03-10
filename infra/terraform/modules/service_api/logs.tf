resource "aws_cloudwatch_log_group" "api" {
  name              = "tzlink-api"
  retention_in_days = 7
}