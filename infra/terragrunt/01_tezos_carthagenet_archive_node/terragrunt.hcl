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
  TZ_MODE    = "archive"

  MIN_INSTANCE_NUMBER = 1
  DESIRED_INSTANCE_NUMBER = 1
  MAX_INSTANCE_NUMBER = 5

  HEALTH_CHECK_GRACE_PERIOD = 1620 #sec (=27min)

  CPU_OUT_SCALING_COOLDOWN = 1800 #sec (=30min)
  CPU_OUT_SCALING_THRESHOLD = 40 #%
  CPU_OUT_EVALUATION_PERIODS = 27 #min

  CPU_DOWN_SCALING_COOLDOWN = 300 #sec
  CPU_DOWN_SCALING_THRESHOLD = 5 #%
  CPU_DOWN_EVALUATION_PERIODS = 5 #min

  INSTANCE_TYPE = "i3.large"
  KEY_PAIR_NAME = "adbo"
}