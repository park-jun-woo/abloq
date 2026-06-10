variable "project" {
  description = "Project name used as resource prefix, e.g. my-blog"
  type        = string
}

variable "aws_region" {
  description = "AWS region for the S3 origin bucket, e.g. ap-northeast-2"
  type        = string
}

variable "domain" {
  description = "Site domain served by CloudFront, e.g. blog.example.com"
  type        = string
}

variable "enable_waf" {
  description = "Attach a WAFv2 web ACL (managed common rules + rate limit) to CloudFront"
  type        = bool
  default     = false
}

variable "waf_rate_limit" {
  description = "WAF rate limit (requests per 5 minutes per IP)"
  type        = number
  default     = 2000
}

variable "enable_logs" {
  description = "Create a CloudFront standard-log bucket (abloq visibility pipeline input)"
  type        = bool
  default     = false
}

variable "log_retention_days" {
  description = "Days to keep CloudFront logs"
  type        = number
  default     = 90
}

variable "indexnow_key" {
  description = "IndexNow key; served at /{key}.txt when non-empty"
  type        = string
  default     = ""
}
