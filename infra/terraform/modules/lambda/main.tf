data "aws_iam_role" "tzlink_lambdas_access" {
  name = "tzlink_lambdas_access"
}

resource "aws_s3_bucket" "lambda_dedicated" {
  bucket = var.S3_BUCKET_NAME
  acl    = "private"

  tags = {
    Name      = var.S3_BUCKET_NAME
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_lambda_function" "executor" {
  s3_bucket     = aws_s3_bucket.lambda_dedicated.bucket
  s3_key        = var.S3_CODE_PATH
  function_name = var.LAMBDA_NAME
  role          = data.aws_iam_role.tzlink_lambdas_access.arn
  handler       = "main"
  runtime       = "go1.x"
  description   = var.LAMBDA_DESCRIPTION
  timeout       = 900 # sec

  vpc_config {
      subnet_ids = var.LAMBDA_VPC_CONFIG_ENABLE ? tolist(data.aws_subnet_ids.lambda[0].ids) : []
      security_group_ids = var.LAMBDA_VPC_CONFIG_ENABLE ? [ data.aws_security_group.lambda[0].id ] : []
  }

  environment {
    variables = var.LAMBDA_ENVIRONMENT_VARIABLES
  }

  tags = {
    Name      = var.LAMBDA_NAME
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.executor.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.cron.arn
}

