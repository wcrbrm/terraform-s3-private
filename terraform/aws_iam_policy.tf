resource "aws_iam_policy" "kms_policy" {
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": [
      "kms:Encrypt",
      "kms:Decrypt",
      "kms:DescribeKey"
    ],
    "Resource": "${aws_kms_key.storage.arn}"
  }
}
EOF
}

resource "aws_iam_policy" "storage_policy" {
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": [
      "s3:*"
    ],
    "Resource": [
      "${aws_s3_bucket.storage.arn}"
    ]
  }
}
EOF
}

