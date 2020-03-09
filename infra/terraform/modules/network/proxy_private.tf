################### proxy A ######################

resource "aws_subnet" "private_proxy_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_PROXY_CIDR, 1, 0)
  availability_zone = "${var.REGION}a"

  map_public_ip_on_launch = false
  tags = {
    Name        = format("tzlink-%s-private-proxy-a", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route" "private_proxy_to_gateway_a" {
  route_table_id         = aws_route_table.tzlink_private_proxy_a.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id             = aws_nat_gateway.public_proxy_a.id
}

resource "aws_route_table" "tzlink_private_proxy_a" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name        = format("tzlink-%s-private-proxy-a", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }

  # route {
  #   cidr_block = "0.0.0.0/0"
  #   gateway_id = aws_nat_gateway.public_proxy_a.id
  # }
}

resource "aws_route_table_association" "private_to_proxy_a" {
  subnet_id      = aws_subnet.private_proxy_a.id
  route_table_id = aws_route_table.tzlink_private_proxy_a.id
}

################### proxy B ######################

resource "aws_subnet" "private_proxy_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PRIVATE_PROXY_CIDR, 1, 1)
  availability_zone = "${var.REGION}b"

  map_public_ip_on_launch = false
  tags = {
    Name        = format("tzlink-%s-private-proxy-b", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route" "private_proxy_to_gateway_b" {
  route_table_id         = aws_route_table.tzlink_private_proxy_b.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id             = aws_nat_gateway.public_proxy_b.id
}

resource "aws_route_table" "tzlink_private_proxy_b" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name        = format("tzlink-%s-private-proxy-b", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }

  # route {
  #   cidr_block = "0.0.0.0/0"
  #   gateway_id = aws_nat_gateway.public_proxy_b.id
  # }
}

resource "aws_route_table_association" "private_to_proxy_b" {
  subnet_id      = aws_subnet.private_proxy_b.id
  route_table_id = aws_route_table.tzlink_private_proxy_b.id
}