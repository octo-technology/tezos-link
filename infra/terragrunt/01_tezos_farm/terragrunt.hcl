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

  MIN_INSTANCE_NUMBER = 2
  DESIRED_INSTANCE_NUMBER = 2
  MAX_INSTANCE_NUMBER = 5

  INSTANCE_TYPE = "i3.large"
  KEY_PAIR_NAME = "adbo"
}