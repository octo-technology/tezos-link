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
  database_master_username = "tezoslink_team"
  database_name            = "tezoslink"
}