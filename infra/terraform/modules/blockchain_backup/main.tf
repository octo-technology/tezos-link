data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
    Name        = format("tzlink-%s", var.ENV)
    Environment = var.ENV
  }
}

resource "aws_s3_bucket" "blockchain_data" {
  bucket = format("tzlink-blockchain-data-%s", var.ENV)
  acl    = "private"

  tags = {
    Name        = format("tzlink-blockchain-data-%s", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_vpc_endpoint" "s3" {
  vpc_id       = data.aws_vpc.tzlink.id
  service_name = "com.amazonaws.${var.REGION}.s3"

  tags = {
    Name        = format("tzlink-blockchain-data-%s", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}