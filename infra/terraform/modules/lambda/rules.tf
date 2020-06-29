resource "aws_cloudwatch_event_rule" "cron" {
  name                = format("tzlink_%s_cronjob", var.LAMBDA_NAME)
  description         = var.LAMBDA_DESCRIPTION
  schedule_expression = var.RUN_EVERY

  tags = {
    Name      = format("tzlink_%s_cronjob", var.LAMBDA_NAME)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_cloudwatch_event_target" "cron" {
  rule      = aws_cloudwatch_event_rule.cron.name
  target_id = aws_lambda_function.executor.function_name
  arn       = aws_lambda_function.executor.arn
}
