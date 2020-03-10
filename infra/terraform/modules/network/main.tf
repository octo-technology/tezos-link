resource "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name      = "tzlink"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}
