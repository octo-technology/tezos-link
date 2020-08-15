include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/s3_shared_bucket"
}

inputs = {
  s3_bucket_name = "tzlink-metric-cleaner-lambda"
}
