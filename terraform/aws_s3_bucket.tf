resource "aws_s3_bucket" "storage" {
  bucket = "${var.application}-private-${random_string.postfix.result}"
  acl    = "private"
  region = var.aws_region
  force_destroy = var.force_destroy

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = aws_kms_key.storage.arn
        sse_algorithm     = "aws:kms"
      }
    }
  }

  tags = var.tags
}

resource "aws_s3_bucket_public_access_block" "storage_access_block" {
  bucket = aws_s3_bucket.storage.id

  # PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
  # PUT Object calls will fail if the request includes an object ACL.
  block_public_acls       = true
  # Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
  block_public_policy     = true
  # Ignore public ACLs on this bucket and any objects that it contains.
  ignore_public_acls      = true
  # Only the bucket owner and AWS Services can access this buckets if it has a public policy.
  restrict_public_buckets = false
}
