resource "aws_subnet" "tzlink_farm_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_FARM_CIDR, 1, 0)
  availability_zone = "${var.REGION}a"

  map_public_ip_on_launch = true
  tags = {
    Name        = format("tzlink-%s-farm", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route_table_association" "tzlink_public_to_farm_a" {
  subnet_id      = aws_subnet.tzlink_farm_a.id
  route_table_id = aws_route_table.tzlink_public.id
}

resource "aws_subnet" "tzlink_farm_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_FARM_CIDR, 1, 1)
  availability_zone = "${var.REGION}b"

  map_public_ip_on_launch = true
  tags = {
    Name        = format("tzlink-%s-farm", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route_table_association" "tzlink_public_to_farm_b" {
  subnet_id      = aws_subnet.tzlink_farm_b.id
  route_table_id = aws_route_table.tzlink_public.id
}