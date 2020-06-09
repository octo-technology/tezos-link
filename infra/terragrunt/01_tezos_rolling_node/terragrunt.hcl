include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network"]
}

terraform {
  source = "../../terraform/modules/tezos_node"
}

inputs = {
  TZ_NETWORK = "mainnet"
  TZ_MODE    = "rolling"

  MIN_INSTANCE_NUMBER = 1
  DESIRED_INSTANCE_NUMBER = 1
  MAX_INSTANCE_NUMBER = 5

  INSTANCE_TYPE = "t2.small"
  KEY_PAIR_NAME = "adbo"
}