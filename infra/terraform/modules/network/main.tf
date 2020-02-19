resource "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name        = format("tzlink-%s", var.ENV)
    Project     = "tezos-link"
    Environment = var.ENV
    BuildWith   = "terraform"
  }
}
