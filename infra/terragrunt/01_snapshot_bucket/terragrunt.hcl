include {
  path = "${find_in_parent_folders()}"
}

dependencies {
  paths = ["../00_network"]
}

terraform {
  source = "../../terraform/modules/s3_shared_bucket"
}

inputs = {
  s3_bucket_name = "tzlink-blockchain-data"
  vpc_endpoint_enabled = true
  route_table_name = "tzlink-public"
}