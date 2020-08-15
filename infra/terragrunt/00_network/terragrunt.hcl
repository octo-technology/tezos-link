include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/network"
}

inputs = {
  subnet_tz_farm_cidr             = "10.1.0.0/24"
  subnet_tz_public_ecs_cidr       = "10.1.1.0/24"
  subnet_tz_private_ecs_cidr      = "10.1.2.0/24"
  subnet_tz_private_database_cidr = "10.1.3.0/24"
}