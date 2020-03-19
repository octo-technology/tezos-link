include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network"]
}

terraform {
  source = "../../terraform/modules/rds_cluster"
}

inputs = {
  DATABASE_PASSWORD = "${get_env("TF_VAR_DATABASE_PASSWORD", "xxxx")}"
}