resource "aws_autoscaling_policy" "cpu_out" {
  name                   = format("tzlink-%s-%s-out-cpu", var.TZ_NETWORK, var.TZ_MODE)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = 420 #sec (7 min)
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "cpu_down" {
  name                   = format("tzlink-%s-%s-down-cpu", var.TZ_NETWORK, var.TZ_MODE)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = 300
  policy_type            = "SimpleScaling"
}


resource "aws_cloudwatch_metric_alarm" "cpu_scale_out" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-%s-out-cpu", var.TZ_NETWORK, var.TZ_MODE)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 40 #%

  period             = 60
  evaluation_periods = 2

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization >= 40% (duration >= 1min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_out.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_down" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-%s-down-cpu", var.TZ_NETWORK, var.TZ_MODE)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = 5 #%

  period             = 60
  evaluation_periods = 30

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization <= 5% (duration >= 5min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_down.arn]
}