data "aws_iam_policy_document" "tzlink_trusted_entity" {
    statement {
        sid = "1"

        actions = [ "sts:AssumeRole" ]

        principals {
            type = "Service"
            identifiers = ["ec2.amazonaws.com"]
        }

        effect = "Allow"
    }
}

data "aws_iam_policy_document" "tzlink_backup_access" {
  statement {
    sid = "1"

    actions = [
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:PutObject",
      "s3:ListBucket",
    ]

    not_resources = [
      "arn:aws:s3:::tzlink-tfstate",
      "arn:aws:s3:::tzlink-tfstate/*"
    ]
  }
}


data "aws_iam_policy_document" "tzlink_backup_personal_access" {
  statement {
    sid = "1"

    actions = [
      "iam:*InstanceProfile*",
      "s3:Get*Role*",
      "s3:List*Role*",
      "s3:PassRole",
    ]

    resources = [
      "arn:aws:iam:::role/tzlink_backup_access", # ARN role que je n'ai pas
    ]
  }
}