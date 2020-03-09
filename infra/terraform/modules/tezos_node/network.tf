data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name        = "tzlink-dev"
  }
}

data "aws_subnet_ids" "tzlink" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = "tzlink-dev-farm-*"
  }
}