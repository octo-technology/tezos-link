[
  {
    "name": "${task_name}",
    "image": "${task_image}",
    "portMappings": [
      {
        "containerPort": ${task_port},
        "hostPort": ${task_port},
        "protocol": "tcp"
      }
    ],
    "cpu": ${task_cpu},
    "memory": ${task_memory},
    "environment": [
        {
          "name": "SERVER_PORT",
          "value": "${task_port}"
        },
        {
          "name": "DATABASE_URL",
          "value": "${database_url}"
        },
        {
          "name": "DATABASE_USERNAME",
          "value": "${database_username}"
        },
        {
          "name": "DATABASE_TABLE",
          "value": "${database_name}"
        },
        {
          "name": "ENV",
          "value": "${configuration_file}"
        }
    ],
    "secrets": [
      {
        "name": "DATABASE_PASSWORD",
        "valueFrom": "${database_password_arn}"
      }
    ],
    "logConfiguration": {
      "logDriver": "awslogs",
      "options": {
        "awslogs-group": "${log_group_name}",
        "awslogs-region": "${log_group_region}",
        "awslogs-stream-prefix": "${log_group_stream_prefix}"
      }
    }
  }
]