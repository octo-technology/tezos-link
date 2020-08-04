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
  TZ_NETWORK = "carthagenet"
  TZ_MODE    = "rolling"

  MIN_INSTANCE_NUMBER = 2
  DESIRED_INSTANCE_NUMBER = 2
  MAX_INSTANCE_NUMBER = 6

  HEALTH_CHECK_GRACE_PERIOD = 900 #sec (=15min)

  CPU_OUT_SCALING_COOLDOWN = 300 #sec
  CPU_OUT_SCALING_THRESHOLD = 30 #%
  CPU_OUT_EVALUATION_PERIODS = 5 #min

  CPU_DOWN_SCALING_COOLDOWN = 300 #sec
  CPU_DOWN_SCALING_THRESHOLD = 5 #%
  CPU_DOWN_EVALUATION_PERIODS = 5 #min

  RESPONSETIME_OUT_SCALING_COOLDOWN = 60 #sec
  RESPONSETIME_OUT_SCALING_THRESHOLD = 0.6 #sec
  RESPONSETIME_OUT_EVALUATION_PERIODS = 2 #min

  INSTANCE_TYPE = "t3.small"
  KEY_PAIR_NAME = "adbo"
}