include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network"]
}

terraform {
  source = "../../terraform/modules/blockchain_backup"
}

inputs = {
  ENABLE_FLOW_LOGS = true
}