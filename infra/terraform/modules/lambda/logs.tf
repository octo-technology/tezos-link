resource "aws_cloudwatch_log_group" "snapshot" {
  name              = "tzlink-snapshot"
  retention_in_days = 7
}
