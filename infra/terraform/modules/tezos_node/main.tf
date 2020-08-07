terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

# Network configuration

data "aws_vpc" "tzlink" {
  cidr_block = var.vpc_cidr

  tags = {
    Name = "tzlink"
  }
}

data "aws_subnet_ids" "tzlink" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-farm-*"
    Project = var.project_name
  }
}

# Network access control

resource "aws_security_group" "tezos_node_lb" {
  name        = format("tezos_farm_lb_%s_%s", var.tz_network, var.tz_mode)
  description = format("Security group for tezos loadbalancer targeting the %s network", var.tz_network)
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "http_ingress_for_loadbalancer_from_vpc" {
  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.tezos_node_lb.id
}

resource "aws_security_group_rule" "all_egress_inside_vpc_for_loadbalancer" {
  type        = "egress"
  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.tezos_node_lb.id
}

resource "aws_security_group" "tezos_node" {
  name        = format("tezos_farm_%s_%s", var.tz_network, var.tz_mode)
  description = format("Security group for tezos nodes in network %s", var.tz_network)
  vpc_id      = data.aws_vpc.tzlink.id

  tags = {
    Name    = format("tzlink_farm_%s_%s", var.tz_network, var.tz_mode)
    Project = var.project_name
  }
}

resource "aws_security_group_rule" "rpc_ingress_for_tezos_node" {
  type                     = "ingress"
  from_port                = 8000
  to_port                  = 8000
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.tezos_node_lb.id

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "p2p_ingress_for_tezos_mainnet_node" {
  count       = var.tz_network == "carthagenet" ? 0 : 1
  type        = "ingress"
  from_port   = 9732
  to_port     = 9732
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]


  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "p2p_ingress_for_tezos_carthagenet_node" {
  count       = var.tz_network == "carthagenet" ? 1 : 0
  type        = "ingress"
  from_port   = 19732
  to_port     = 19732
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "ssh_ingress_for_tezos_node_from_vpc" {
  count       = var.tz_mode == "archive" ? 1 : 0
  type        = "ingress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "rpc_egress_for_tezos_node" {
  type        = "egress"
  from_port   = 8732
  to_port     = 8732
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "websock_egress_for_tezos_node" {
  type        = "egress"
  from_port   = 9732
  to_port     = 9732
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "http_egress_for_tezos_node" {
  type        = "egress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "https_egress_for_tezos_node" {
  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "ssh_egress_for_tezos_node_from_vpc" {
  count       = var.tz_mode == "archive" ? 1 : 0
  type        = "egress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.tezos_node.id
}

# Internal DNS record configuration

data "aws_route53_zone" "tezoslink_private" {
  name         = "internal.tezoslink.io."
  private_zone = true
}

resource "aws_route53_record" "internal_lb_farm" {
  zone_id = data.aws_route53_zone.tezoslink_private.zone_id
  name    = format("%s-%s.internal.tezoslink.io", var.tz_network, var.tz_mode)
  type    = "A"

  alias {
    name                   = aws_alb.tz_farm.dns_name
    zone_id                = aws_alb.tz_farm.zone_id
    evaluate_target_health = false
  }
}

# Instance pool configuration

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_launch_configuration" "tz_node" {
  instance_type = var.instance_type
  image_id      = data.aws_ami.ubuntu.id

  # Manually created role based on iam/ec2-assume-role
  iam_instance_profile = "tzlink_ec2"

  key_name = var.key_pair_name

  security_groups = [aws_security_group.tezos_node.id]

  enable_monitoring = true

  user_data = templatefile("${path.module}/templates/user_data_${var.tz_mode}.tpl", {
    network           = var.tz_network
    lambda_public_key = var.lambda_public_key
  })

  root_block_device {
    volume_type           = "io1"
    volume_size           = 10 #Gb
    encrypted             = true
    delete_on_termination = true
    iops                  = 500 # can be only 50 * volume_size
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "tz_nodes" {
  name = format("tzlink-%s-%s", var.tz_network, var.tz_mode)

  desired_capacity = var.desired_instance_number
  max_size         = var.max_instance_number
  min_size         = var.min_instance_number

  health_check_grace_period = var.health_check_grace_period
  health_check_type         = "ELB"
  force_delete              = true
  launch_configuration      = aws_launch_configuration.tz_node.id
  vpc_zone_identifier       = tolist(data.aws_subnet_ids.tzlink.ids)

  default_cooldown = 180 #sec (= 3min)

  enabled_metrics = ["GroupInServiceInstances", "GroupPendingInstances", "GroupStandbyInstances", "GroupTerminatingInstances", "GroupTotalInstances"]

  target_group_arns = [aws_alb_target_group.tz_farm.arn]

  lifecycle {
    create_before_destroy = true
  }

  # Permits to suicide instance and handle Accenture policy 56
  # This will create a new fresh instance every 10 days and remove
  # the old one.
  max_instance_lifetime = 864000 #sec (= 10days)

  tags = [
    {
      key                 = "Name"
      value               = format("tzlink-%s", var.tz_network)
      propagate_at_launch = true
    },
    {
      key                 = "Project"
      value               = var.project_name
      propagate_at_launch = true
    },
    {
      key                 = "Mode"
      value               = var.tz_mode
      propagate_at_launch = true
    },
    {
      key                 = "Accessibility"
      value               = "public"
      propagate_at_launch = true
    }
  ]
}

# Loadbalancer configuration

resource "aws_alb" "tz_farm" {
  name            = format("tzlink-farm-%s-%s", var.tz_network, var.tz_mode)
  subnets         = tolist(data.aws_subnet_ids.tzlink.ids)
  security_groups = [aws_security_group.tezos_node_lb.id]
  internal        = true

  tags = {
    Name    = format("tzlink-farm-%s-%s", var.tz_network, var.tz_mode)
    Project = var.project_name
  }
}

resource "aws_alb_target_group" "tz_farm" {
  name        = format("tzlink-farm-%s-%s", var.tz_network, var.tz_mode)
  port        = 8000
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.tzlink.id
  target_type = "instance"

  health_check {
    enabled             = true
    path                = "/chains/main/blocks/head"
    port                = 8000
    protocol            = "HTTP"
    healthy_threshold   = 2
    unhealthy_threshold = 2
    interval            = 30
    timeout             = 29
  }

  tags = {
    Name    = format("tzlink-farm-%s-%s", var.tz_network, var.tz_mode)
    Project = var.project_name
  }

  depends_on = [aws_alb.tz_farm]
}

resource "aws_alb_listener" "tz_farm_http" {
  load_balancer_arn = aws_alb.tz_farm.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.tz_farm.arn
    type             = "forward"
  }

  depends_on = [aws_alb_target_group.tz_farm]
}

# Autoscaling system

resource "aws_autoscaling_policy" "cpu_out" {
  name                   = format("tzlink-%s-%s-out-cpu", var.tz_network, var.tz_mode)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = var.cpu_out_scaling_cooldown
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "cpu_down" {
  name                   = format("tzlink-%s-%s-down-cpu", var.tz_network, var.tz_mode)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "-1"
  cooldown               = var.cpu_down_scaling_cooldown
  policy_type            = "SimpleScaling"
}

resource "aws_autoscaling_policy" "response_time_out" {
  name                   = format("tzlink-%s-%s-out-response-time", var.tz_network, var.tz_mode)
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.name
  adjustment_type        = "ChangeInCapacity"
  scaling_adjustment     = "1"
  cooldown               = var.responsetime_out_scaling_cooldown
  policy_type            = "SimpleScaling"
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_out" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-%s-out-cpu", var.tz_network, var.tz_mode)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = var.cpu_out_scaling_threshold

  period             = 60 # sec
  evaluation_periods = var.cpu_out_evaluation_periods

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization >= ${var.cpu_out_scaling_threshold}% (duration >= ${var.cpu_out_evaluation_periods}min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_out.arn]
}

resource "aws_cloudwatch_metric_alarm" "cpu_scale_down" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-%s-down-cpu", var.tz_network, var.tz_mode)
  namespace           = "AWS/EC2"
  statistic           = "Average"
  metric_name         = "CPUUtilization"
  comparison_operator = "LessThanOrEqualToThreshold"
  threshold           = var.cpu_down_scaling_threshold

  period             = 60 # sec
  evaluation_periods = var.cpu_down_evaluation_periods

  dimensions = {
    AutoScalingGroupName = aws_autoscaling_group.tz_nodes.name
  }

  alarm_description = "Average CPUUtilization <= ${var.cpu_down_scaling_threshold}% (duration >= ${var.cpu_down_evaluation_periods}min)"
  alarm_actions     = [aws_autoscaling_policy.cpu_down.arn]
}

resource "aws_cloudwatch_metric_alarm" "response_time_scale_out" {
  actions_enabled = true

  alarm_name          = format("tzlink-%s-%s-out-response_time", var.tz_network, var.tz_mode)
  namespace           = "AWS/ApplicationELB"
  statistic           = "Maximum"
  metric_name         = "TargetResponseTime"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  threshold           = var.responsetime_out_scaling_threshold

  period             = 60 #sec
  evaluation_periods = var.responsetime_out_evaluation_periods

  dimensions = {
    TargetGroup  = aws_alb_target_group.tz_farm.arn_suffix
    LoadBalancer = aws_alb.tz_farm.arn_suffix
  }

  treat_missing_data = "notBreaching"

  alarm_description = "Maximum ResponseTime >= ${var.responsetime_out_scaling_threshold}% (duration >= ${var.responsetime_out_evaluation_periods}min)"
  alarm_actions     = [aws_autoscaling_policy.response_time_out.arn]
}