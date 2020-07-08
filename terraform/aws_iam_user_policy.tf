
resource "aws_iam_user_policy" "user_storage_policy" {
  name = "${aws_s3_bucket.storage.bucket}-user-storage-${aws_iam_user.user.name}"
  user = aws_iam_user.user.name

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Action": [ "s3:*" ],
        "Effect": "Allow",
        "Resource": "${aws_s3_bucket.storage.arn}"
    },
    {
        "Action": [ "s3:*" ],
        "Effect": "Allow",
        "Resource": "${aws_s3_bucket.storage.arn}/*"
    },
    {
        "Effect": "Allow",
        "Action": [ "kms:*" ],
        "Resource": "${aws_kms_key.storage.arn}"
    }
  ]
}
EOF
}