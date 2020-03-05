resource "aws_ecs_task_definition" "proxy" {
  family                   = "proxy"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.PROXY_CPU
  memory                   = var.PROXY_MEMORY

  container_definitions = <<EOF
[
  {
    "cpu": var.PROXY_CPU,
    "image": var.PROXY_DOCKER_IMAGE,
    "memory": var.PROXY_MEMORY
    "name": "proxy",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": var.PROXY_PORT,
        "hostPort": var.PROXY_PORT
      }
    ]
  }
]
EOF
}

resource "aws_ecs_service" "proxy" {
  name            = "proxy"
  cluster         = aws_ecs_cluster.proxy.id
  task_definition = aws_ecs_task_definition.proxy.arn
  desired_count   = var.PROXY_DESIRED_COUNT
  launch_type     = "FARGATE"

  network_configuration {
    security_groups = [aws_security_group.proxy_ecs_task.id]
    subnets         = [data.aws_subnet_ids.tzlink_proxy.id]
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.proxy.id
    container_name   = "proxy"
    container_port   = var.PROXY_PORT
  }

  depends_on = [aws_alb_listener.proxy]
}
