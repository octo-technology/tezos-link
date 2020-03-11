data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name = "tzlink"
  }
}

data "aws_subnet_ids" "database" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name = "tzlink-private-database-*"
  }
}