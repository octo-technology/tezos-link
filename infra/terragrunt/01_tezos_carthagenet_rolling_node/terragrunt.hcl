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
  tz_network = "carthagenet"
  tz_mode    = "rolling"

  min_instance_number     = 2
  desired_instance_number = 2
  max_instance_number     = 6

  health_check_grace_period = 900 #sec (=15min)

  cpu_out_scaling_cooldown   = 300 #sec (=5min)
  cpu_out_scaling_threshold  = 40 #%
  cpu_out_evaluation_periods = 5 #min

  cpu_down_scaling_cooldown   = 300 #sec
  cpu_down_scaling_threshold  = 5 #%
  cpu_down_evaluation_periods = 5 #min

  responsetime_out_scaling_cooldown   = 60 #sec
  responsetime_out_scaling_threshold  = 0.6 #sec
  responsetime_out_evaluation_periods = 2 #min

  instance_type = "t3.small"
  key_pair_name = "adbo" # TODO : use a shared keypair
}