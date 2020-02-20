data "aws_ami" "centos" {
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

data "aws_vpc" "tzlink" {
  cidr_block = var.VPC_CIDR

  tags = {
      Name = format("tzlink-%s", var.ENV)
      Environment = var.ENV
  }
}

data "aws_subnet_ids" "tzlink" {
  vpc_id = data.aws_vpc.tzlink.id

  tags = {
      Name = format("tzlink-%s-farm-*", var.ENV)
  }
}

resource "aws_security_group" "tezos_node" {
  name        = "tezos_node"
  description = "Tezos node security group"
  vpc_id      = data.aws_vpc.tzlink.id
}

resource "aws_security_group_rule" "ssh_ingress" {
  type            = "ingress"
  from_port       = 22
  to_port         = 22
  protocol        = "tcp"
  cidr_blocks     = ["0.0.0.0/0"]

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_security_group_rule" "all_egress" {
  type            = "egress"
  from_port       = 0
  to_port         = 0
  protocol        = "-1"
  cidr_blocks     = ["0.0.0.0/0"]
  # prefix_list_ids = ["pl-12c4e678"] # TODO: useful for VPC endpoint. To check later.

  security_group_id = aws_security_group.tezos_node.id
}

resource "aws_instance" "tz_node" {
  ami           = data.aws_ami.centos.id
  instance_type = "t2.micro"
  subnet_id = tolist(data.aws_subnet_ids.tzlink.ids)[0]

  #key_name = "adbo"

  associate_public_ip_address = true

  vpc_security_group_ids = [ aws_security_group.tezos_node.id ]

  tags = {
    Name        = format("tzlink-%s-test", var.ENV)
    Project     = "tezos-link"
    Environment = var.ENV
    BuildWith   = "terraform"
    Trigramme   = "adbo"
  }
}