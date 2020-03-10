resource "aws_internet_gateway" "tzlink" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name        = "tzlink"
    Project     = var.PROJECT_NAME
    Environment = var.ENV
    BuildWith   = var.BUILD_WITH
  }
}

resource "aws_route_table" "tzlink_public" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name      = "tzlink-public"
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_route" "public_route" {
  route_table_id         = aws_route_table.tzlink_public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.tzlink.id
}