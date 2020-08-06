remote_state {
  backend = "s3"
  config = {
      bucket = "tzlink-remote-state"
      key    = "${get_env("TF_VAR_ENV", "prod")}/${path_relative_to_include()}/terraform.tfstate"
      encrypt = true
      region  = "eu-west-1"
      dynamodb_table = "tzlink-remote-state-lock"
  }
}

terraform {
  extra_arguments "custom_vars" {
    commands  = get_terraform_commands_that_need_vars()

    required_var_files = [
      "${get_parent_terragrunt_dir()}/../common.tfvars"
    ]
  }
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF

provider "aws" {
  version = "~> 2.0"
  region = var.region
}

EOF
}