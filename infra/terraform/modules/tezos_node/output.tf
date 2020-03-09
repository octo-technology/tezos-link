output "lb_farm_endpoint" {
  value = aws_elb.tz_farm.dns_name
}