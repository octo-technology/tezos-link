resource "aws_autoscaling_policy" "out" {
  name                   = format("tzlink-%s-out", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = "60"
  policy_type            = "SimpleScaling"
}

resource "aws_cloudwatch_metric_alarm" "out" {
  alarm_name          = format("tzlink-%s-out", var.TZ_NETWORK)
  namespace           = "AWS/ELB"
  statistic           = "Average"
  metric_name         = "RequestCount"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = "200"

  period             = "60"
  evaluation_periods = "1"

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average RequestCount >=200 req/sec (duration >=1min)"
  alarm_actions     = [aws_autoscaling_policy.out.arn]
}

resource "aws_autoscaling_policy" "down" {
  name                   = format("tzlink-%s-down", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = "300"
  policy_type            = "SimpleScaling"
}

resource "aws_cloudwatch_metric_alarm" "down" {
  alarm_name          = format("tzlink-%s-down", var.TZ_NETWORK)
  namespace           = "AWS/ELB"
  statistic           = "Average"
  metric_name         = "RequestCount"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = "50"

  period             = "60"
  evaluation_periods = "5"

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average RequestCount <=50 req/sec (duration >=1min)"
  alarm_actions     = [aws_autoscaling_policy.down.arn]
}