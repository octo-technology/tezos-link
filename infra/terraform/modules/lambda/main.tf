resource "aws_s3_bucket" "snapshot_lambda" {
  bucket = format("tzlink-snapshot-lambda-%s", var.ENV)
  acl    = "private"

  tags = {
    Name      = format("tzlink-snapshot-lambda-%s", var.ENV)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "random_id" "zipfile" {
  keepers = {
    version = sha256(file("${path.module}/../../../cmd/snapshot"))
  }

  byte_length = 8
}

resource "archive_file" "snapshot_lambda" {
  type        = "zip"
  source_dir  = "${path.module}/../../../bin/snapshot"
  output_path = "${path.module}/../../../bin/tzlink-snapshot-lambda-${random_id.zipfile.hex}.zip"
  depends_on  = ["random_id.zipfile"]
}

resource "aws_s3_bucket_object" "snapshot_lambda" {
  key         = "tzlink-snapshot-lambda.zip"
  bucket      = aws_s3_bucket.snapshot_lambda.bucket
  source      = "${path.module}/../../../bin/tzlink-snapshot-lambda-${random_id.zipfile.hex}.zip"
  etag        = sha256(file("${path.module}/../../../bin/tzlink-snapshot-lambda-${random_id.zipfile.hex}.zip"))
  depends_on  = ["archive_file.lambda", "aws_s3_bucket.snapshot_lambda"]
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "snapshot_lambda" {
  s3_bucket        = aws_s3_bucket.snapshot_lambda.bucket
  s3_key           = aws_s3_bucket.snapshot_lambda
  function_name    = "snapshot"
  role             = aws_iam_role.iam_for_lambda.arn
  handler          = "snapshot"
  source_code_hash = filebase64sha256("snapshot_lambda.zip")
  runtime          = "go1.x"
  description      = "Snapshot exporter Lambda"

  environment {
    variables = {
      NODE_IP = "0.0.0.0"
    }
  }

  depends_on = ["aws_s3_bucket_object.snapshot_lambda", "random_id.zipfile", "archive_file.snapshot_lambda"]
}
