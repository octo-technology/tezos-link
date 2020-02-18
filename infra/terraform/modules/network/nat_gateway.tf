resource "aws_eip" "tzlink_nat" {
  vpc = true
}

resource "aws_nat_gateway" "tzlink" {
  allocation_id = aws_eip.tzlink_nat.id
  subnet_id     = aws_subnet.tzlink_farm.id
}