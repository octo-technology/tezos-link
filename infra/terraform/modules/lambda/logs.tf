resource "aws_cloudwatch_log_group" "snapshot" {
  name              = "/aws/lambda/snapshot"
  retention_in_days = 7
}
