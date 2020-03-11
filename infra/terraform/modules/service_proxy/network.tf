data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name = "tzlink"
  }
}

data "aws_subnet_ids" "tzlink_public_proxy" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = "tzlink-public-ecs-*"
  }
}

data "aws_subnet_ids" "tzlink_private_proxy" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = "tzlink-private-ecs-*"
  }
}