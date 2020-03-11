resource "aws_subnet" "private_database_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_DATABASE_CIDR, 1, 0)
  availability_zone = "${var.REGION}a"

  map_public_ip_on_launch = false
  tags = {
    Name      = "tzlink-private-database-a"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_subnet" "private_database_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_DATABASE_CIDR, 1, 1)
  availability_zone = "${var.REGION}b"

  map_public_ip_on_launch = false
  tags = {
    Name      = "tzlink-private-database-b"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}