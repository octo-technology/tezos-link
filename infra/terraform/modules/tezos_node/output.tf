output "lb_farm_endpoint" {
  value = aws_alb.tz_farm.dns_name
}
