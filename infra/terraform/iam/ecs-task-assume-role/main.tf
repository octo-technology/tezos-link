data "aws_iam_policy_document" "tzlink_ecs_task_trusted_entity" {
  statement {

    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }

    effect = "Allow"
  }
}

data "aws_iam_policy_document" "tzlink_ecs_task_ecr_and_awslogs" {
  statement {

    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]

    resources = [
      "*"
    ]
  }

  statement {

    actions = [
      "secretsmanager:GetSecretValue"
    ]

    resources = [
      "arn:aws:secretsmanager:eu-west-1:*:secret:tzlink-database-password"
    ]
  }
}