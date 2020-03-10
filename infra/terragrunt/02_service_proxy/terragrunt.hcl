include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../01_ecs_cluster", "../01_rds_cluster"]
}

terraform {
  source = "../../terraform/modules/service_proxy"
}

inputs = {
  PROXY_DOCKER_IMAGE_NAME = "louptheronlth/tezos-link"
  PROXY_DOCKER_IMAGE_VERSION = "proxy-dev"
  PROXY_DESIRED_COUNT = 1

  PROXY_CONFIGURATION_FILE = "local"
  PROXY_PORT = 8001
  PROXY_CPU = 256
  PROXY_MEMORY = 512

  TEZOS_FARM_PORT = 80
}