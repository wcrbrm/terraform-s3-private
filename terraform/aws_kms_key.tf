resource "aws_kms_key" "storage" {
  description = "Used to unseal s3 bucket data"
  tags = var.tags
}
