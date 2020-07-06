resource "aws_iam_user" "user" {
  name = "${var.application}-${random_string.postfix.result}"
  path = "/app/${var.application}/"
  force_destroy = true
  tags = var.tags
}

resource "aws_iam_access_key" "key" {
  user = aws_iam_user.user.name
}

