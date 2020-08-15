output "trusted_entity" {
  value = data.aws_iam_policy_document.tzlink_lambda_trusted_entity.json
}

output "tzlink_global_lambda_access" {
  value = data.aws_iam_policy_document.tzlink_global_lambda_access.json
}

output "tzlink_snapshot_lambda_access" {
  value = data.aws_iam_policy_document.tzlink_snapshot_lambda_access.json
}

output "tzlink_metric_lambda_access" {
  value = data.aws_iam_policy_document.tzlink_metric_lambda_access.json
}

output "tzlink_secretmanager_lambda_access" {
  value = data.aws_iam_policy_document.tzlink_secretmanager_lambda_access.json
}