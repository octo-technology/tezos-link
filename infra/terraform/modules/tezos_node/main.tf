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

  user_data = templatefile("${path.module}/user_data.tpl", {
    network           = var.TZ_NETWORK
    lambda_public_key = file("${path.module}/lambda_public_key")
    computed_network  = var.TZ_NETWORK == "mainnet" ? "babylonnet" : var.TZ_NETWORK
  })

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "tz_nodes" {
  name = format("tzlink-%s", var.TZ_NETWORK)

  desired_capacity = var.DESIRED_INSTANCE_NUMBER
  max_size         = var.MAX_INSTANCE_NUMBER
  min_size         = var.MIN_INSTANCE_NUMBER

  health_check_grace_period = 1800 # 30mins
  health_check_type         = "ELB"
  force_delete              = true
  launch_configuration      = aws_launch_configuration.tz_node.id
  vpc_zone_identifier       = tolist(data.aws_subnet_ids.tzlink.ids)

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
    }
  ]
}

resource "aws_autoscaling_attachment" "tz_farm" {
  autoscaling_group_name = aws_autoscaling_group.tz_nodes.id
  alb_target_group_arn   = aws_alb_target_group.tz_farm.arn
}


resource "aws_alb" "tz_farm" {
  name            = format("tzlink-farm-%s", var.TZ_NETWORK)
  subnets         = tolist(data.aws_subnet_ids.tzlink.ids)
  security_groups = [aws_security_group.tezos_node_lb.id]
  internal        = true

  tags = {
    Name      = format("tzlink-farm-%s", var.TZ_NETWORK)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}

resource "aws_alb_target_group" "tz_farm" {
  name        = format("tzlink-farm-%s", var.TZ_NETWORK)
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
    timeout             = 3
  }

  tags = {
    Name      = format("tzlink-farm-%s", var.TZ_NETWORK)
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
