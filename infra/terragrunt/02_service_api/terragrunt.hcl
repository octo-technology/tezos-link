include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = [
            "../00_ecs_cluster",
            "../01_rds_cluster"
          ]
}

terraform {
  source = "../../terraform/modules/service_api"
}

inputs = {
  docker_image_name        = "louptheronlth/tezos-link"
  docker_image_version     = "${get_env("TF_VAR_DOCKER_IMAGE_VERSION", "proxy-api")}"
  desired_container_number = 1

  port   = 8001
  cpu    = 256
  memory = 512

  configuration_file = "local"

  database_master_username = "tezoslink_team"
  database_name            = "tezoslink"
}