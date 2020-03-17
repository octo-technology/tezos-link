resource "aws_cloudwatch_event_rule" "every_twelve_hours" {
  name                = "tzlink_snapshot_export_cronjob"
  description         = "Send a snapshot export request"
  schedule_expression = "rate(12 hours)"

  tags = {
    Name      = "tzlink-snapshot-rule"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_cloudwatch_event_target" "snapshot_every_twelve_hours" {
  rule      = aws_cloudwatch_event_rule.every_twelve_hours.name
  target_id = aws_lambda_function.snapshot_lambda.function_name
  arn       = aws_lambda_function.snapshot_lambda.arn
}
