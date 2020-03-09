include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../01_ecs_cluster", "../01_rds_cluster"]
}

terraform {
  source = "../../terraform/modules/service_backend"
}

inputs = {
  BACKEND_DOCKER_IMAGE_NAME = "louptheronlth/tezos-link"
  BACKEND_DOCKER_IMAGE_VERSION = "api-dev"
  BACKEND_DESIRED_COUNT = 1

  BACKEND_CONFIGURATION_FILE = "local"

  BACKEND_PORT = 80
  BACKEND_CPU = 256
  BACKEND_MEMORY = 512
}