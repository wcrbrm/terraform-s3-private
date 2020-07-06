resource "aws_s3_bucket" "storage" {
  bucket = "${var.application}-private-${random_string.postfix.result}"
  acl    = "private"
  region = var.aws_region

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = aws_kms_key.storage.arn
        sse_algorithm     = "aws:kms"
      }
    }
  }

  tags = map(
    "kubernetes.io/cluster/${var.application}", "owned",
  )
}
