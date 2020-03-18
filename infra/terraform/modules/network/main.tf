resource "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name      = "tzlink"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}
