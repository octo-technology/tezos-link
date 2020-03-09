include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network"]
}

terraform {
  source = "../../terraform/modules/ecs_cluster"
}

inputs = {
  ENABLE_FLOW_LOGS = true
}