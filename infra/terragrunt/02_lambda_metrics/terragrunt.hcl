include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  S3_BUCKET_NAME = "tzlink-metrics-lambda-${get_env("TF_VAR_ENV", "dev")}"
  S3_CODE_PATH = "v1.0.0/metrics.zip"

  LAMBDA_PURPOSE = "metrics"
  LAMBDA_DESCRIPTION = "RDS old metrics cleaner lambda"
  LAMBDA_ENVIRONMENT_VARIABLES = {
    RDS_ENDPOINT = "postgres://xxxxx:xxxxx@xxxxx"
  }

  RUN_EVERY = "rate(24 hours)"
}
