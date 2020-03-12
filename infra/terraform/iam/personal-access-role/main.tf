data "aws_iam_policy_document" "tzlink_personal_access" {
  statement {

    actions = [
      "iam:Get*Role*",
      "iam:List*Role*",
      "iam:PassRole",
    ]

    resources = [
      "arn:aws:iam::609827314188:role/tzlink_backup_access",
      "arn:aws:iam::609827314188:role/tzlink_ecs_tasks_access",
      "arn:aws:iam::609827314188:role/tzlink_lambda_access",
    ]
  }

  statement {

    actions = [
      "iam:*InstanceProfile*",
    ]

    resources = [
      "*",
    ]
  }
}
