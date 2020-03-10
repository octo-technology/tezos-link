resource "aws_cloudwatch_log_group" "backend" {
  name              = "tzlink-backend"
  retention_in_days = 7
}