resource "aws_internet_gateway" "tzlink" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name        = format("tzlink-%s", var.ENV)
    Project     = "tezos-link"
    Environment = "all"
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}