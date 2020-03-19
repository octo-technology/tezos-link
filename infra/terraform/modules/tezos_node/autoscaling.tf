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

resource "aws_autoscaling_policy" "requestcountbytarget_out" {
  name                   = format("tzlink-%s-out-requestcountbytarget", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = 1200
  policy_type            = "SimpleScaling"
}


########################

resource "aws_cloudwatch_metric_alarm" "latency_scale_out" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-out-latency", var.TZ_NETWORK)
  namespace           = "AWS/ApplicationELB"
  statistic           = "Maximum"
  metric_name         = "TargetResponseTime"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 5 # sec

  period             = 60
  evaluation_periods = 1

  dimensions = {
    Loadbalancer = aws_alb.tz_farm.arn_suffix
    TargetGroup  = aws_alb_target_group.tz_farm.arn_suffix
  }

  alarm_description = "Maximum Latency >= 5sec (duration >= 1min)"
  alarm_actions     = [aws_autoscaling_policy.latency_out.arn]
}

resource "aws_cloudwatch_metric_alarm" "requestcountbytarget_scale_out" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-out-requestcountbytarget", var.TZ_NETWORK)
  namespace           = "AWS/ApplicationELB"
  statistic           = "Sum"
  metric_name         = "RequestCountPerTarget"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 120 # requests/sec

  period             = 60
  evaluation_periods = 1

  dimensions = {
    Loadbalancer = aws_alb.tz_farm.arn_suffix
    TargetGroup  = aws_alb_target_group.tz_farm.arn_suffix
  }

  alarm_description = "Average TargetResponseTime >= 120 req/sec (duration >= 1min)"
  alarm_actions     = [aws_autoscaling_policy.requestcountbytarget_out.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_out" {
  actions_enabled = false

  alarm_name          = format("tzlink-%s-out-cpu", var.TZ_NETWORK)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = 40 #%

  period             = 60
  evaluation_periods = 1

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization >= 40% (duration >= 1min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_out.arn]
}

##################

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

resource "aws_autoscaling_policy" "requestcountbytarget_down" {
  name                   = format("tzlink-%s-down-requestcountbytarget", var.TZ_NETWORK)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = 300
  policy_type            = "SimpleScaling"
}

##################

resource "aws_cloudwatch_metric_alarm" "latency_scale_down" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-down-latency", var.TZ_NETWORK)
  namespace           = "AWS/ApplicationELB"
  statistic           = "Maximum"
  metric_name         = "TargetResponseTime"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = 2

  period             = 60
  evaluation_periods = 5

  dimensions = {
    Loadbalancer = aws_alb.tz_farm.arn_suffix
    TargetGroup  = aws_alb_target_group.tz_farm.arn_suffix
  }

  alarm_description = "Average Latency <= 2sec (duration >= 5min)"
  alarm_actions     = [aws_autoscaling_policy.latency_down.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_down" {
  actions_enabled = false

  alarm_name          = format("tzlink-%s-down-cpu", var.TZ_NETWORK)
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

resource "aws_cloudwatch_metric_alarm" "requestcountbytarget_scale_down" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-down-requestcountbytarget", var.TZ_NETWORK)
  namespace           = "AWS/ApplicationELB"
  statistic           = "Sum"
  metric_name         = "RequestCountPerTarget"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = 20 # requests/sec

  period             = 60
  evaluation_periods = 5

  dimensions = {
    Loadbalancer = aws_alb.tz_farm.arn_suffix
    TargetGroup  = aws_alb_target_group.tz_farm.arn_suffix
  }

  alarm_description = "Average CPUUtilization <= 20 req/sec (duration >= 5min)"
  alarm_actions     = [aws_autoscaling_policy.requestcountbytarget_down.arn]
}