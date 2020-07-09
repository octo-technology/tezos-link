include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  S3_BUCKET_NAME = "tzlink-snapshot-lambda-${get_env("TF_VAR_ENV", "dev")}"
  S3_CODE_PATH = "v1.0.0/snapshot.zip"

  LAMBDA_NAME = "snapshot-mainnet"
  LAMBDA_DESCRIPTION = "Snapshot exporter lambda for mainnet"
  LAMBDA_ENVIRONMENT_VARIABLES = {
    NODE_USER = "ec2-user"
    S3_REGION = "eu-west-1"
    S3_BUCKET = "tzlink-snapshot-lambda-${get_env("TF_VAR_ENV", "dev")}"
    S3_LAMBDA_KEY = "snapshot_lambda_key"
    NETWORK = "mainnet"
  }

  LAMBDA_VPC_CONFIG_ENABLE = false

  RUN_EVERY = "cron(0 1 * * ? *)"
}
