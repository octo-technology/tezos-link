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

  HEALTH_CHECK_GRACE_PERIOD = 900 #sec (=15min)

  CPU_OUT_SCALING_COOLDOWN = 1080 #sec
  CPU_OUT_SCALING_THRESHOLD = 40 #%
  CPU_OUT_EVALUATION_PERIODS = 12 #min

  CPU_DOWN_SCALING_COOLDOWN = 300 #sec
  CPU_DOWN_SCALING_THRESHOLD = 5 #%
  CPU_DOWN_EVALUATION_PERIODS = 5 #min

  INSTANCE_TYPE = "t3.small"
  KEY_PAIR_NAME = "adbo"
}