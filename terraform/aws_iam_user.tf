resource "aws_iam_user" "user" {
  name = "${var.application}-${random_string.postfix.result}"
}

resource "aws_iam_access_key" "key" {
  user = aws_iam_user.user.name
}

