resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = format("/aws/lambda/%s", var.LAMBDA_NAME)
  retention_in_days = 7
}
