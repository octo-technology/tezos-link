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
  tz_network = "mainnet"
  tz_mode    = "archive"

  min_instance_number     = 1
  desired_instance_number = 1
  max_instance_number     = 5

  health_check_grace_period = 1620 #sec (=27min)

  cpu_out_scaling_cooldown   = 1800 #sec (=30min)
  cpu_out_scaling_threshold  = 40 #%
  cpu_out_evaluation_periods = 30 #min

  cpu_down_scaling_cooldown   = 300 #sec
  cpu_down_scaling_threshold  = 5 #%
  cpu_down_evaluation_periods = 5 #min

  responsetime_out_scaling_cooldown   = 60 #sec
  responsetime_out_scaling_threshold  = 1.0 #sec
  responsetime_out_evaluation_periods = 2 #min

  instance_type = "i3.large"
  key_pair_name = "adbo" # TODO : use a shared keypair
}