resource "aws_subnet" "private_database_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_DATABASE_CIDR, 1, 0)
  availability_zone = "${var.REGION}a"

  map_public_ip_on_launch = false
  tags = {
    Name        = format("tzlink-%s-private-database-a", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_subnet" "private_database_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_DATABASE_CIDR, 1, 1)
  availability_zone = "${var.REGION}b"

  map_public_ip_on_launch = false
  tags = {
    Name        = format("tzlink-%s-private-database-b", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}