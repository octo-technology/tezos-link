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
    DATABASE_USERNAME = "tezoslink_team"
    DATABASE_NAME     = "tezoslink"
    DATABASE_URL      = "tzlink-database.cluster-cjwwybxbgnpm.eu-west-1.rds.amazonaws.com"
    DATABASE_PASSWORD_SECRET_ARN = "arn:aws:secretsmanager:eu-west-1:912174778846:secret:tzlink-database-password-IwehdC"
  }
  run_every = "cron(0 1 * * ? *)"

  # based on 00_lambda_snapshot_bucket
  bucket_name = "tzlink-metric-cleaner-lambda"
  code_path   = "v1.0.0/metrics.zip"

  # network configuration (if needed)
  vpc_config_enable = true
  subnet_name = "tzlink-private-database-*"
  security_group_name = "tzlink-database"
}
