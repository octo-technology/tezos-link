output "trusted_entity" {
  value = data.aws_iam_policy_document.tzlink_trusted_entity.json
}

output "ec2_instance_s3_snapshot_access" {
  value = data.aws_iam_policy_document.ec2_instance_s3_snapshot_access.json
}

output "ec2_instance_ssm_profile" {
  value = data.aws_iam_policy_document.ec2_instance_ssm_profile.json
}