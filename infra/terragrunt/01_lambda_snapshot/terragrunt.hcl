include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  S3_BUCKET_NAME = "tzlink-snapshot-lambda-${get_env("TF_VAR_ENV", "dev")}"
  S3_CODE_PATH = "v1.0.0/snapshot.zip"

  LAMBDA_PURPOSE = "snapshot"
  LAMBDA_DESCRIPTION = "Snapshot exporter lambda"
  LAMBDA_ENVIRONMENT_VARIABLES = {
    NODE_USER = "ec2-user"
    S3_REGION = "eu-west-1"
    S3_BUCKET = "tzlink-snapshot-lambda-${get_env("TF_VAR_ENV", "dev")}"
    S3_LAMBDA_KEY = "snapshot_lambda_key"
  }

  RUN_EVERY = "rate(12 hours)"
}
