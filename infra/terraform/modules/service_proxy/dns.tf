data "aws_route53_zone" "tezoslink" {
  name = "tezoslink.io."
}

resource "aws_route53_record" "mainnet" {
  zone_id = data.aws_route53_zone.tezoslink.zone_id
  name    = "mainnet.tezoslink.io"
  type    = "A"

  alias {
    name                   = aws_alb.proxy.dns_name
    zone_id                = aws_alb.proxy.zone_id
    evaluate_target_health = false
  }
}