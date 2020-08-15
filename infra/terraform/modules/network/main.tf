terraform {
  required_version = "0.12.20"
  backend "s3" {}
}

# Manually created role based on iam/vpc-flowlog-assume-role
data "aws_iam_role" "tzlink_vpc_flowlog" {
  name = "tzlink_vpc_flowlog"
}

# AWS VPC

resource "aws_vpc" "tzlink" {
  cidr_block = var.vpc_cidr

  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name    = "tzlink"
    Project = var.project_name
  }
}

resource "aws_default_security_group" "tzlink" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name    = format("%s-default", var.project_name)
    Project = var.project_name
  }
}

# VPC Logflow

resource "aws_cloudwatch_log_group" "tzlink_vpc_flowlog" {
  name = "tzlink_vpc_flowlog"

  tags = {
    Name    = format("%s_vpc_flowlog", var.project_name)
    Project = var.project_name
  }
}

resource "aws_cloudwatch_log_stream" "tzlink_vpc_flowlog" {
  name           = "tzlink_vpc_flowlog"
  log_group_name = aws_cloudwatch_log_group.tzlink_vpc_flowlog.name
}

resource "aws_flow_log" "tzlink_flowlog" {
  vpc_id          = aws_vpc.tzlink.id
  log_destination = aws_cloudwatch_log_group.tzlink_vpc_flowlog.arn
  traffic_type    = "ALL"

  iam_role_arn = data.aws_iam_role.tzlink_vpc_flowlog.arn

  tags = {
    Name    = format("%s_vpc_flowlog", var.project_name)
    Project = var.project_name
  }
}

# DNS Zones

resource "aws_route53_zone" "tzlink_public" {
  name = "tezoslink.io"
}

resource "aws_route53_zone" "tzlink_private" {
  name = "internal.tezoslink.io"

  vpc {
    vpc_id = aws_vpc.tzlink.id
  }
}

# Internet gateway connected to the public subnets

resource "aws_internet_gateway" "tzlink" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink"
    Project = var.project_name
  }
}

resource "aws_route_table" "tzlink_public" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-public"
    Project = var.project_name
  }
}

resource "aws_route" "public_route" {
  route_table_id         = aws_route_table.tzlink_public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.tzlink.id
}

# Public subnet for Tezos-Nodes A

resource "aws_subnet" "public_farm_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_farm_cidr, 1, 0)
  availability_zone = "${var.region}a"

  map_public_ip_on_launch = true
  tags = {
    Name    = "tzlink-farm-a"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "public_to_farm_a" {
  subnet_id      = aws_subnet.public_farm_a.id
  route_table_id = aws_route_table.tzlink_public.id
}

# Public subnet for Tezos-Nodes B

resource "aws_subnet" "public_farm_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_farm_cidr, 1, 1)
  availability_zone = "${var.region}b"

  map_public_ip_on_launch = true
  tags = {
    Name    = "tzlink-farm-b"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "public_to_farm_b" {
  subnet_id      = aws_subnet.public_farm_b.id
  route_table_id = aws_route_table.tzlink_public.id
}

# Public subnet for ECS's loadbalancer A

resource "aws_subnet" "public_ecs_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_public_ecs_cidr, 1, 0)
  availability_zone = "${var.region}a"

  map_public_ip_on_launch = true
  tags = {
    Name    = "tzlink-public-ecs-a"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "public_to_ecs_a" {
  subnet_id      = aws_subnet.public_ecs_a.id
  route_table_id = aws_route_table.tzlink_public.id
}

resource "aws_eip" "gateway_public_ecs_a" {
  vpc        = true
  depends_on = [aws_internet_gateway.tzlink]
}

resource "aws_nat_gateway" "public_ecs_a" {
  subnet_id     = aws_subnet.public_ecs_a.id
  allocation_id = aws_eip.gateway_public_ecs_a.id
}

# Public subnet for ECS's loadbalancer B

resource "aws_subnet" "public_ecs_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_public_ecs_cidr, 1, 1)
  availability_zone = "${var.region}b"

  map_public_ip_on_launch = true
  tags = {
    Name    = "tzlink-public-ecs-b"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "public_to_ecs_b" {
  subnet_id      = aws_subnet.public_ecs_b.id
  route_table_id = aws_route_table.tzlink_public.id
}

resource "aws_eip" "gateway_public_ecs_b" {
  vpc        = true
  depends_on = [aws_internet_gateway.tzlink]
}

resource "aws_nat_gateway" "public_ecs_b" {
  subnet_id     = aws_subnet.public_ecs_b.id
  allocation_id = aws_eip.gateway_public_ecs_b.id
}

# Private subnet for ECS A

resource "aws_subnet" "private_ecs_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_private_ecs_cidr, 1, 0)
  availability_zone = "${var.region}a"

  map_public_ip_on_launch = false
  tags = {
    Name    = "tzlink-private-ecs-a"
    Project = var.project_name
  }
}

resource "aws_route" "private_ecs_to_gateway_a" {
  route_table_id         = aws_route_table.tzlink_private_ecs_a.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.public_ecs_a.id
}

resource "aws_route_table" "tzlink_private_ecs_a" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-private-ecs-a"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "private_to_ecs_a" {
  subnet_id      = aws_subnet.private_ecs_a.id
  route_table_id = aws_route_table.tzlink_private_ecs_a.id
}

# Private subnet for ECS B

resource "aws_subnet" "private_ecs_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_private_ecs_cidr, 1, 1)
  availability_zone = "${var.region}b"

  map_public_ip_on_launch = false
  tags = {
    Name    = "tzlink-private-ecs-b"
    Project = var.project_name
  }
}

resource "aws_route" "private_ecs_to_gateway_b" {
  route_table_id         = aws_route_table.tzlink_private_ecs_b.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.public_ecs_b.id
}

resource "aws_route_table" "tzlink_private_ecs_b" {
  vpc_id = aws_vpc.tzlink.id

  tags = {
    Name    = "tzlink-private-ecs-b"
    Project = var.project_name
  }
}

resource "aws_route_table_association" "private_to_ecs_b" {
  subnet_id      = aws_subnet.private_ecs_b.id
  route_table_id = aws_route_table.tzlink_private_ecs_b.id
}

# Private subnet for Database A

resource "aws_subnet" "private_database_a" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_private_database_cidr, 1, 0)
  availability_zone = "${var.region}a"

  map_public_ip_on_launch = false
  tags = {
    Name    = "tzlink-private-database-a"
    Project = var.project_name
  }
}

# Private subnet for Database B

resource "aws_subnet" "private_database_b" {
  vpc_id            = aws_vpc.tzlink.id
  cidr_block        = cidrsubnet(var.subnet_tz_private_database_cidr, 1, 1)
  availability_zone = "${var.region}b"

  map_public_ip_on_launch = false
  tags = {
    Name    = "tzlink-private-database-b"
    Project = var.project_name
  }
}

# SecurityGroup for AWS_API

resource "aws_security_group" "aws_api" {
  name        = "aws_api_endpoint"
  description = "Security group applyied to vpc_endpoint to access them by http/https"
  vpc_id      = aws_vpc.tzlink.id

  tags = {
    Name    = "aws_api_endpoint"
    Project = "tezos-link"
  }
}

resource "aws_security_group_rule" "http_ingress_for_aws_api_from_vpc" {
  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.aws_api.id
}

resource "aws_security_group_rule" "https_ingress_for_aws_api_from_vpc" {
  type        = "ingress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = [var.vpc_cidr]

  security_group_id = aws_security_group.aws_api.id
}

# S3 VPC Endpoint

resource "aws_vpc_endpoint" "s3" {
  vpc_id       = aws_vpc.tzlink.id
  service_name = format("com.amazonaws.%s.s3", var.region)

  route_table_ids = [aws_route_table.tzlink_public.id]

  tags = {
    Name    = "tzlink-public-s3"
    Project = var.project_name
  }
}

# EC2 VPC Endpoint

resource "aws_vpc_endpoint" "ec2" {
  vpc_id            = aws_vpc.tzlink.id
  service_name      = format("com.amazonaws.%s.ec2", var.region)
  vpc_endpoint_type = "Interface"

  subnet_ids = [
    aws_subnet.public_farm_a.id,
    aws_subnet.public_farm_b.id
  ]

  security_group_ids = [
    aws_security_group.aws_api.id
  ]

  private_dns_enabled = true

  tags = {
    Name    = "tzlink-public-ec2"
    Project = var.project_name
  }
}

# SecretsManager VPC Endpoint

resource "aws_vpc_endpoint" "secretsmanager" {
  vpc_id            = aws_vpc.tzlink.id
  service_name      = format("com.amazonaws.%s.secretsmanager", var.region)
  vpc_endpoint_type = "Interface"

  subnet_ids = [
    aws_subnet.public_farm_a.id,
    aws_subnet.public_farm_b.id
  ]

  security_group_ids = [
    aws_security_group.aws_api.id
  ]

  private_dns_enabled = true

  tags = {
    Name    = "tzlink-public-secretsmanager"
    Project = var.project_name
  }
}