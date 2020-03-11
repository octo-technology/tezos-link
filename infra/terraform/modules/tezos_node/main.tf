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

  health_check_grace_period = 300
  health_check_type         = "EC2"
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
  elb                    = aws_elb.tz_farm.id
}

resource "aws_elb" "tz_farm" {
  name            = format("tzlink-farm-%s", var.TZ_NETWORK)
  subnets         = tolist(data.aws_subnet_ids.tzlink.ids)
  internal        = true
  security_groups = [aws_security_group.tezos_node_lb.id]

  listener {
    instance_port     = 8000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }

  health_check {
    healthy_threshold   = 2
    unhealthy_threshold = 2
    timeout             = 3
    target              = "HTTP:8000/chains/main/blocks/head"
    interval            = 30
  }

  cross_zone_load_balancing   = true
  idle_timeout                = 400
  connection_draining         = true
  connection_draining_timeout = 400

  tags = {
    Name      = format("tzlink-farm-%s", var.TZ_NETWORK)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
}
