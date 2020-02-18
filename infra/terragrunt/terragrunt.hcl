remote_state {
  backend = "s3"
  config = {
      bucket = "tezos-link-tfstate"
      key    = "${get_env("TF_VAR_ENV", "dev")}/${path_relative_to_include()}/terraform.tfstate"
      encrypt = true
      region  = "eu-west-1"
  }
}

terraform {
  extra_arguments "custom_vars" {
    commands  = get_terraform_commands_that_need_vars()

    required_var_files = [
      "${get_parent_terragrunt_dir()}/../${get_env("TF_VAR_ENV", "dev")}.tfvars",
      "${get_parent_terragrunt_dir()}/../common.tfvars"
    ]
  }
}