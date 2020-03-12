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

data "aws_iam_policy_document" "tzlink_lambda_access" {
  statement {

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "s3:*"
    ]

    resources = [
      "arn:aws:s3:::tzlink-blockchain-data-dev",
      "arn:aws:s3:::tzlink-blockchain-data-dev/*",
      "arn:aws:s3:::tzlink-snapshot-lambda-dev",
      "arn:aws:s3:::tzlink-snapshot-lambda-dev/*"
    ]
  }
}
