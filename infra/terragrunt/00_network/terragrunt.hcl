include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/network"
}

inputs = {
  ENABLE_FLOW_LOGS = true
}