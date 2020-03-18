resource "aws_autoscaling_policy" "latency_out" {
  name                   = format("tzlink-%s-out-latency", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = 1200
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "cpu_out" {
  name                   = format("tzlink-%s-out-cpu", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = 1200
  policy_type            = "SimpleScaling"
}

resource "aws_cloudwatch_metric_alarm" "latency_scale_out" {
  alarm_name = format("tzlink-%s-out-latency", var.TZ_NETWORK)
  namespace  = "AWS/ELB"
  statistic  = "Average"
  # Can be swapped with "RequestCount" if needed
  metric_name         = "Latency"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 5 # sec

  period             = 60
  evaluation_periods = 1

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average Latency >= 5sec (duration >= 1min)" # "Average RequestCount >=200 req/sec (duration >=1min)"
  alarm_actions     = [aws_autoscaling_policy.latency_out.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_out" {
  alarm_name = format("tzlink-%s-out-cpu", var.TZ_NETWORK)
  namespace  = "AWS/EC2"
  statistic  = "Average"
  # Can be swapped with "RequestCount" if needed
  metric_name         = "CPUUtilization"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 40 #%

  period             = 60
  evaluation_periods = 1

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization >= 40% (duration >= 1min)" # "Average RequestCount >=200 req/sec (duration >=1min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_out.arn]
}

resource "aws_autoscaling_policy" "latency_down" {
  name                   = format("tzlink-%s-down-latency", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = 300
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "cpu_down" {
  name                   = format("tzlink-%s-down-cpu", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = 300
  policy_type            = "SimpleScaling"
}

resource "aws_cloudwatch_metric_alarm" "latency_scale_down" {
  alarm_name = format("tzlink-%s-down-latency", var.TZ_NETWORK)
  namespace  = "AWS/ELB"
  # Can be swapped with "RequestCount" if needed
  statistic           = "Average"
  metric_name         = "Latency"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = 2

  period             = 60
  evaluation_periods = 5

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average Latency <= 2sec (duration >= 5min)" # "Average RequestCount <=50 req/sec (duration >=1min)"
  alarm_actions     = [aws_autoscaling_policy.latency_down.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_down" {
  alarm_name          = format("tzlink-%s-out-down", var.TZ_NETWORK)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = 5 #%

  period             = 60
  evaluation_periods = 5

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization <= 5% (duration >= 5min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_down.arn]
}