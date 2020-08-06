include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network", "../01_rds_cluster"]
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  # lambda configuration
  name        = "metrics"
  description = "RDS metrics cleaner"
  environment_variables = {
    DATABASE_USERNAME = "administrator"
    DATABASE_TABLE    = "tezoslink"
    DATABASE_URL      = "TODO" # TODO : to change
    DATABASE_PASSWORD = "TODO" #"${get_env("TF_VAR_DATABASE_PASSWORD")}" # TODO : find a way to retrieve the RDS password in AWS SSM
  }
  run_every = "cron(0 1 * * ? *)"

  # based on 00_lambda_snapshot_bucket
  bucket_name = "tzlink-metric-cleaner-lambda"
  code_path   = "v1.0.0/metrics.zip"

  # network configuration (if needed)
  vpc_config_enable = true
  subnet_name = "tzlink-private-database-*"
  security_group = "database"
}
