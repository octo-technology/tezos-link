data "aws_vpc" "tzlink" {
  count = var.LAMBDA_VPC_CONFIG_ENABLE ? 1 : 0

  cidr_block = var.VPC_CIDR

  tags = {
    Name = "tzlink"
  }
}

data "aws_subnet_ids" "lambda" {
  count = var.LAMBDA_VPC_CONFIG_ENABLE ? 1 : 0

  vpc_id = data.aws_vpc.tzlink[0].id

  tags = {
    Name = var.LAMBDA_SUBNET_NAME
  }
}

data "aws_security_group" "lambda" {
  count = var.LAMBDA_VPC_CONFIG_ENABLE ? 1 : 0
  vpc_id = data.aws_vpc.tzlink[0].id

  tags = {
    Name        = var.LAMBDA_SECURITY_GROUP_NAME
    Project     = var.PROJECT_NAME
  }
}
