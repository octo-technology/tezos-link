output "tzlink_vpc_flowlog" {
  value = data.aws_iam_policy_document.tzlink_vpc_flowlog.json
}

output "trusted_entity" {
  value = data.aws_iam_policy_document.tzlink_trusted_entity.json
}