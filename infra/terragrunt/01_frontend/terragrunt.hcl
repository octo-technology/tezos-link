include {
  path = "${find_in_parent_folders()}"
}

terraform {
  source = "../../terraform/modules/cloudfront"
}

inputs = {
  certificate_arn = "arn:aws:acm:us-east-1:912174778846:certificate/6f9b1bac-5177-488a-a350-95b10d056a2a"
}