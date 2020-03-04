data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name        = format("tzlink-%s", var.ENV)
    Environment = var.ENV
  }
}

data "aws_subnet_ids" "tzlink_proxy" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = format("tzlink-%s-proxy-*", var.ENV)
  }
}