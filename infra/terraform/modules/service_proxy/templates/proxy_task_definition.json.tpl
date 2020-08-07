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
    "ulimits": [
      {
        "name": "nofile",
        "softLimit": 10000,
        "hardLimit": 12000
      }
    ],
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
          "name": "ARCHIVE_NODES_URL",
          "value": "${tezos_archive_hostname}"
        },
        {
          "name": "TEZOS_ARCHIVE_PORT",
          "value": "${tezos_archive_port}"
        },
        {
          "name": "ROLLING_NODES_URL",
          "value": "${tezos_rolling_hostname}"
        },
        {
          "name": "TEZOS_ROLLING_PORT",
          "value": "${tezos_rolling_port}"
        },
        {
          "name": "TEZOS_NETWORK",
          "value": "${tezos_network}"
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