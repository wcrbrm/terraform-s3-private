output "bucket_arn" {
  value       = aws_s3_bucket.storage.arn
}

output "bucket_id" {
  value       = aws_s3_bucket.storage.id
  description = "Bucket Name (aka ID)"
}

output "bucket_region" {
  value = aws_s3_bucket.storage.region
}

output "access_key" {
  value = aws_iam_access_key.key.id
  sensitive = true
}

output "secret_key" {
  value = aws_iam_access_key.key.secret
  sensitive = true
}
