variable "project" {
  description = "Project name used as resource prefix"
  type        = string
}

variable "retention_days" {
  description = "Days to keep CloudFront logs"
  type        = number
  default     = 90
}
