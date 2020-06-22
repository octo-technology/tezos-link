resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = format("/aws/lambda/%s", var.LAMBDA_PURPOSE)
  retention_in_days = 7
}
