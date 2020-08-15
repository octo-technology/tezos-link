include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = [
            "../00_ecs_cluster",
            "../01_rds_cluster",
            "../01_tezos_carthagenet_archive_node",
            "../01_tezos_carthagenet_rolling_node"
          ]
}

terraform {
  source = "../../terraform/modules/service_proxy"
}

inputs = {
  docker_image_name        = "louptheronlth/tezos-link"
  docker_image_version     = "${get_env("TF_VAR_DOCKER_IMAGE_VERSION", "proxy-dev")}"
  desired_container_number = 2

  port   = 8001
  cpu    = 256
  memory = 512

  configuration_file = "prod"
  tz_network         = "carthagenet"

  database_master_username = "tezoslink_team"
  database_name            = "tezoslink"

  farm_archive_port = 80
  farm_rolling_port = 80
}
