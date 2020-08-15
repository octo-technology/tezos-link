include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = [
            "../00_lambda_snapshot_bucket",
            "../01_tezos_carthagenet_archive_node"
          ]
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  # lambda configuration
  name        = "snapshot-carthagenet"
  description = "Snapshot exporter lambda for carthagenet"
  environment_variables = {
    NODE_USER = "ubuntu"
    S3_REGION = "eu-west-1"
    S3_BUCKET = "tzlink-snapshot-lambda"
    S3_LAMBDA_KEY = "snapshot_lambda_key"
    NETWORK = "carthagenet"
  }
  run_every = "cron(0 1/12 * * ? *)"

  # based on 00_lambda_snapshot_bucket
  bucket_name = "tzlink-snapshot-lambda"
  code_path   = "v1.0.0/snapshot.zip"

  # network configuration (if needed)
  vpc_config_enable = true
  subnet_name = "tzlink-farm-*"
  security_group_name = "tzlink_farm_carthagenet_archive"
}
