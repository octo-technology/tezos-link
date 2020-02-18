resource "aws_route_table" "tzlink_public" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name        = format("tzlink-%s-public", var.ENV)
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}

resource "aws_route" "public_route" {
  route_table_id = aws_route_table.tzlink_public.id
  gateway_id     = aws_internet_gateway.tzlink.id
}

resource "aws_route_table_association" "tzlink_public_to_farm" {
  subnet_id      = aws_subnet.tzlink_farm.id
  route_table_id = aws_route_table.tzlink_public.id
}