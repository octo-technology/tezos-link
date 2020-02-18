resource "aws_subnet" "tzlink_farm" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = var.SUBNET_TZ_FARM_CIDR
  availability_zone = var.REGION

  map_public_ip_on_launch = true
  tags {
    Name        = format("tzlink-%s-farm", var.ENV)
    Project     = "tezos-link"
    Environment = var.ENV
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}

resource "aws_route_table_association" "tzlink_public_to_farm" {
  subnet_id      = aws_subnet.tzlink_farm.id
  route_table_id = aws_route_table.tzlink_public.id
}