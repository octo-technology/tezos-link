resource "aws_route53_zone" "tzlink" {
  name = "tezoslink.io"
}

resource "aws_route53_zone" "tzlink_private" {
  name = "internal.tezoslink.io"

  vpc {
    vpc_id = aws_vpc.tzlink.id
  }
}