include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/lambda"
}

inputs = {
  ENABLE_FLOW_LOGS = true
}
