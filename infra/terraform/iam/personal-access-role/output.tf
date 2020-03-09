output "personal_access" {
  value = data.aws_iam_policy_document.tzlink_personal_access.json
}