data "aws_iam_policy_document" "tzlink_lambda_trusted_entity" {
  statement {

    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    effect = "Allow"
  }
}

data "aws_iam_policy_document" "tzlink_global_lambda_access" {
  statement {

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]

    resources = [
      "*"
    ]
  }
}

data "aws_iam_policy_document" "tzlink_snapshot_lambda_access" {
  statement {

    actions = [
      "s3:*"
    ]

    resources = [
      "arn:aws:s3:::tzlink-blockchain-data-dev",
      "arn:aws:s3:::tzlink-blockchain-data-dev/*",
      "arn:aws:s3:::tzlink-snapshot-lambda-dev",
      "arn:aws:s3:::tzlink-snapshot-lambda-dev/*"
    ]
  }

  statement {
    actions = [
      "ec2:Describe*"
    ]

    resources = [
      "*"
    ]
  }
}

data "aws_iam_policy_document" "tzlink_metric_lambda_access" {
  statement {

    actions = [
      "s3:*"
    ]

    resources = [
      "arn:aws:s3:::tzlink-metric-lambda-dev",
      "arn:aws:s3:::tzlink-metric-lambda-dev/*"
    ]
  }

  statement {

    actions = [
      "ec2:CreateNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteNetworkInterface"
    ]

    resources = [
      "*"
    ]
  }
}

data "aws_iam_policy_document" "tzlink_secretmanager_lambda_access" {
  statement {

    actions = [
      "secretsmanager:GetSecretValue"
    ]

    resources = [
      "arn:aws:secretsmanager:eu-west-1:*:secret:tzlink-database-password"
    ]
  }
}