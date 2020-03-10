data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name        = format("tzlink-%s", var.ENV)
    Environment = var.ENV
  }
}

data "aws_subnet_ids" "tzlink_public_ecs" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = format("tzlink-%s-public-proxy-*", var.ENV)
  }
}

data "aws_subnet_ids" "tzlink_private_ecs" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = format("tzlink-%s-private-proxy-*", var.ENV)
  }
}