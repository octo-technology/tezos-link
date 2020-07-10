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
  S3_BUCKET_NAME = "tzlink-metrics-lambda-${get_env("TF_VAR_ENV", "dev")}"
  S3_CODE_PATH = "v1.0.0/metrics.zip"

  LAMBDA_NAME = "metrics"
  LAMBDA_DESCRIPTION = "RDS old metrics cleaner lambda"
  LAMBDA_ENVIRONMENT_VARIABLES = {
    DATABASE_USERNAME = "administrator"
    DATABASE_TABLE = "tezoslink"
    DATABASE_PASSWORD = "${get_env("TF_VAR_DATABASE_PASSWORD")}"
    DATABASE_URL = "tzlink-database.cluster-cmeu9dixowfa.eu-west-1.rds.amazonaws.com"
  }

  LAMBDA_VPC_CONFIG_ENABLE = true
  LAMBDA_SUBNET_NAME = "tzlink-private-database-*"
  LAMBDA_SECURITY_GROUP_NAME = "database"

  RUN_EVERY = "cron(0 1 * * ? *)"
}
