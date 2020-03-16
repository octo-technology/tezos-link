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

resource "aws_instance" "tz_node" {
  count         = var.TZ_NODE_NUMBER
  ami           = data.aws_ami.rhel.id
  instance_type = var.INSTANCE_TYPE
  subnet_id     = tolist(data.aws_subnet_ids.tzlink.ids)[count.index % 2]

  key_name = var.KEY_PAIR_NAME

  associate_public_ip_address = true

  iam_instance_profile = "tzlink_backup_access"

  vpc_security_group_ids = [aws_security_group.tezos_node.id]

  monitoring = true

  user_data = templatefile("${path.module}/user_data.tpl", {
    network           = var.TZ_NETWORK
    lambda_public_key = file("${path.module}/lambda_public_key")
    computed_network  = var.TZ_NETWORK == "mainnet" ? "babylonnet" : var.TZ_NETWORK
  })

  tags = {
    Name      = format("tzlink-%s-%d", var.TZ_NETWORK, count.index)
    Project   = var.PROJECT_NAME
    BuildWith = var.BUILD_WITH
  }
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

  instances                   = tolist(aws_instance.tz_node.*.id)
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
