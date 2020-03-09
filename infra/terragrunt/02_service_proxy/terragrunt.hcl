include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../01_ecs_cluster", "../01_rds_cluster"]
}

terraform {
  source = "../../terraform/modules/service_proxy"
}

inputs = {
  ENABLE_FLOW_LOGS = true
}