resource "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags {
    Name        = format("tzlink-%s", var.ENV)
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}

resource "aws_subnet" "tzlink_farm" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = var.SUBNET_TZ_FARM_CIDR
  availability_zone = var.REGION

  map_public_ip_on_launch = true
  tags {
    Name        = format("tzlink-%s-farm", var.ENV)
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}