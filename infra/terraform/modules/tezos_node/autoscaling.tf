resource "aws_autoscaling_policy" "out" {
  name                   = format("tzlink-%s-out", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = "300"
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "down" {
  name                   = format("tzlink-%s-down", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = "600"
  policy_type            = "SimpleScaling"
}