output "trusted_entity" {
  value = data.aws_iam_policy_document.tzlink_trusted_entity.json
}

output "backup_access" {
  value = data.aws_iam_policy_document.tzlink_backup_access.json
}