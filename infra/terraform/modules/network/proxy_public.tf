################### Public A ######################

resource "aws_subnet" "public_proxy_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PUBLIC_proxy_CIDR, 1, 0)
  availability_zone = "${var.REGION}a"

  map_public_ip_on_launch = true
  tags = {
    Name        = format("tzlink-%s-public-proxy-a", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route_table_association" "public_to_proxy_a" {
  subnet_id      = aws_subnet.public_proxy_a.id
  route_table_id = aws_route_table.tzlink_public.id
}

resource "aws_eip" "gateway_public_proxy_a" {
  vpc        = true
  depends_on = [ aws_internet_gateway.tzlink ]
}

resource "aws_nat_gateway" "public_proxy_a" {
  subnet_id     = aws_subnet.public_proxy_a.id
  allocation_id = aws_eip.gateway_public_proxy_a.id
}

################### Public B ######################

resource "aws_subnet" "public_proxy_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.SUBNET_TZ_PUBLIC_proxy_CIDR, 1, 1)
  availability_zone = "${var.REGION}b"

  map_public_ip_on_launch = true
  tags = {
    Name        = format("tzlink-%s-public-proxy-b", var.ENV)
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route_table_association" "public_to_proxy_b" {
  subnet_id      = aws_subnet.public_proxy_b.id
  route_table_id = aws_route_table.tzlink_public.id
}

resource "aws_eip" "gateway_public_proxy_b" {
  vpc        = true
  depends_on = [ aws_internet_gateway.tzlink ]
}

resource "aws_nat_gateway" "public_proxy_b" {
  subnet_id     = aws_subnet.public_proxy_b.id
  allocation_id = aws_eip.gateway_public_proxy_b.id
}