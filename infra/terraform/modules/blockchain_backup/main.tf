data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name = "tzlink"
  }
}

resource "aws_s3_bucket" "blockchain_data" {
  bucket = format("tzlink-blockchain-data-%s", var.ENV)
  acl    = "private"

  tags = {
    Name      = format("tzlink-blockchain-data-%s", var.ENV)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_vpc_endpoint" "s3" {
  vpc_id       = data.aws_vpc.tzlink.id
  service_name = format("com.amazonaws.%s.s3", var.REGION)

  tags = {
    Name      = "tzlink-blockchain-data"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}