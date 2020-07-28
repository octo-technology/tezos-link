data "aws_ami" "rhel" {
  most_recent = true

  filter {
    name   = "name"
    values = ["RHEL-8.1.0_HVM-20191029-x86_64*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["309956199498"] # RedHat
}

resource "aws_launch_configuration" "tz_node" {
  instance_type = var.INSTANCE_TYPE
  image_id      = data.aws_ami.rhel.id

  iam_instance_profile = "tzlink_backup_access"

  key_name = var.KEY_PAIR_NAME

  security_groups = [aws_security_group.tezos_node.id]

  enable_monitoring = true

  user_data = templatefile("${path.module}/user_data_${var.TZ_MODE}.tpl", {
    network           = var.TZ_NETWORK
    lambda_public_key = file("${path.module}/lambda_public_key")
    computed_network  = var.TZ_NETWORK == "mainnet" ? "babylonnet" : var.TZ_NETWORK
    mode              = var.TZ_MODE
  })

  root_block_device {
    volume_type = "io1"
    volume_size = 10 #Gb
    encrypted = true
    delete_on_termination = true
    iops = 500
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "tz_nodes" {
  name = format("tzlink-%s-%s", var.TZ_NETWORK, var.TZ_MODE)

  desired_capacity = var.DESIRED_INSTANCE_NUMBER
  max_size         = var.MAX_INSTANCE_NUMBER
  min_size         = var.MIN_INSTANCE_NUMBER

  health_check_grace_period = var.HEALTH_CHECK_GRACE_PERIOD
  health_check_type         = "ELB"
  force_delete              = true
  launch_configuration      = aws_launch_configuration.tz_node.id
  vpc_zone_identifier       = tolist(data.aws_subnet_ids.tzlink.ids)

  enabled_metrics = ["GroupInServiceInstances", "GroupPendingInstances", "GroupStandbyInstances", "GroupTerminatingInstances", "GroupTotalInstances"]

  target_group_arns = [aws_alb_target_group.tz_farm.arn]

  lifecycle {
    create_before_destroy = true
  }

  tags = [
    {
      key                 = "Name"
      value               = format("tzlink-%s", var.TZ_NETWORK)
      propagate_at_launch = true
    },
    {
      key                 = "Project"
      value               = var.PROJECT_NAME
      propagate_at_launch = true
    },
    {
      key                 = "BuildWith"
      value               = var.BUILD_WITH
      propagate_at_launch = true
    },
    {
      key                 = "Mode"
      value               = var.TZ_MODE
      propagate_at_launch = true
    }
  ]
}


resource "aws_alb" "tz_farm" {
  name            = format("tzlink-farm-%s-%s", var.TZ_NETWORK, var.TZ_MODE)
  subnets         = tolist(data.aws_subnet_ids.tzlink.ids)
  security_groups = [aws_security_group.tezos_node_lb.id]
  internal        = true

  tags = {
    Name      = format("tzlink-farm-%s-%s", var.TZ_NETWORK, var.TZ_MODE)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "tz_farm" {
  name        = format("tzlink-farm-%s-%s", var.TZ_NETWORK, var.TZ_MODE)
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
    Name      = format("tzlink-farm-%s-%s", var.TZ_NETWORK, var.TZ_MODE)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
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
