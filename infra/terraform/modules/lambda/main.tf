terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

# Manually created role based on iam/lambda-assume-role
data "aws_iam_role" "tzlink_lambdas_access" {
  name = "tzlink_lambda"
}

# Network configuration
# (used if vpc_config_enable=true)

data "aws_vpc" "tzlink" {
  count = var.vpc_config_enable ? 1 : 0

  cidr_block = var.vpc_cidr

  tags = {
    Name = "tzlink"
  }
}

data "aws_subnet_ids" "lambda" {
  count = var.vpc_config_enable ? 1 : 0

  vpc_id = data.aws_vpc.tzlink[0].id

  tags = {
    Name    = var.subnet_name
    Project = var.project_name
  }
}

data "aws_security_group" "lambda" {
  count = var.vpc_config_enable ? 1 : 0

  vpc_id = data.aws_vpc.tzlink[0].id

  tags = {
    Name    = var.security_group_name
    Project = var.project_name
  }
}

# Lambda function configuration

data "aws_s3_bucket" "lambda_dedicated" {
  bucket = var.bucket_name
}

resource "aws_lambda_function" "executor" {
  s3_bucket     = data.aws_s3_bucket.lambda_dedicated.bucket
  s3_key        = var.code_path
  function_name = var.name
  role          = data.aws_iam_role.tzlink_lambdas_access.arn
  handler       = "main"
  runtime       = "go1.x"
  description   = var.description
  timeout       = 900 # sec

  vpc_config {
    subnet_ids         = var.vpc_config_enable ? tolist(data.aws_subnet_ids.lambda[0].ids) : []
    security_group_ids = var.vpc_config_enable ? [data.aws_security_group.lambda[0].id] : []
  }

  environment {
    variables = var.environment_variables
  }

  tags = {
    Name    = var.name
    Project = var.project_name
  }
}

# Cloudwatch Alarm

resource "aws_lambda_permission" "allow_cloudwatch_to_call_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.executor.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.cron.arn
}

resource "aws_cloudwatch_event_rule" "cron" {
  name                = format("tzlink_%s_cronjob", var.name)
  description         = var.description
  schedule_expression = var.run_every

  tags = {
    Name    = format("tzlink_%s_cronjob", var.name)
    Project = var.project_name
  }
}

resource "aws_cloudwatch_event_target" "cron" {
  rule      = aws_cloudwatch_event_rule.cron.name
  target_id = aws_lambda_function.executor.function_name
  arn       = aws_lambda_function.executor.arn
}

# Cloudwatch Logs

resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = format("/aws/lambda/%s", var.name)
  retention_in_days = 7
}

