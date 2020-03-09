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
  INSTANCE_TYPE = "i3.large"
  KEY_PAIR_NAME = "adbo"
}