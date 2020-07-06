
resource "aws_iam_policy_attachment" "kms_policy_attachment" {
  name = "${var.application}-kms-${aws_iam_user.user.name}"
  policy_arn = aws_iam_policy.kms_policy.arn
  users = [
    aws_iam_user.user.name,
  ]
}

resource "aws_iam_policy_attachment" "storage_policy_attachment" {
  name = "${var.application}-storage-${aws_iam_user.user.name}"
  policy_arn = aws_iam_policy.storage_policy.arn
  users = [
    aws_iam_user.user.name,
  ]
}
