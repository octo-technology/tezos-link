include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/s3_shared_bucket"
}

inputs = {
  S3_BUCKET_NAME = "tzlink-metrics-lambda-${get_env("TF_VAR_ENV", "dev")}"
}
