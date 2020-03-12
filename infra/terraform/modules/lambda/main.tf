data "aws_iam_role" "tzlink_lambdas_access" {
  name = "tzlink_lambdas_access"
}

resource "aws_s3_bucket" "snapshot_lambda" {
  bucket = format("tzlink-snapshot-lambda-%s", var.ENV)
  acl    = "private"

  tags = {
    Name      = format("tzlink-snapshot-lambda-%s", var.ENV)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_lambda_function" "snapshot_lambda" {
  s3_bucket        = aws_s3_bucket.snapshot_lambda.bucket
  s3_key           = "v1.0.0/snapshot.zip"
  function_name    = "snapshot"
  role             = data.aws_iam_role.tzlink_lambdas_access.arn
  handler          = "snapshot"
  runtime          = "go1.x"
  description      = "Snapshot exporter Lambda"

  environment {
    variables = {
      NODE_IP = "0.0.0.0"
    }
  }
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_snapshot_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.snapshot_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.every_twelve_hours.arn
}
