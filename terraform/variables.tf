variable "aws_region" {
  type = string
  description = "AWS region for deployment"
}

variable "application" {
  type = string
  description = "Application name (with environment)"
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "Additional tags (e.g. `map('BusinessUnit','XYZ')`"
}

variable "force_destroy" {
  type = bool
  default = true

  description = "Whether it is safe to destroy bucket contents"
}
