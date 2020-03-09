output "trusted_entity" {
  value = data.aws_iam_policy_document.tzlink_ecs_task_trusted_entity.json
}

output "ecs_task_access" {
  value = data.aws_iam_policy_document.tzlink_ecs_task_ecr_and_awslogs.json
}
