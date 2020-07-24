include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/ecs_cluster"
}