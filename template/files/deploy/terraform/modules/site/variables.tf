variable "project" {
  description = "Project name used as resource prefix"
  type        = string
}

variable "domain" {
  description = "Site domain served by CloudFront"
  type        = string
}

variable "web_acl_arn" {
  description = "WAFv2 web ACL ARN to attach (empty = none)"
  type        = string
  default     = ""
}

variable "logs_bucket_domain" {
  description = "CloudFront log bucket domain name (empty = no logging)"
  type        = string
  default     = ""
}

variable "indexnow_key" {
  description = "IndexNow key; served at /{key}.txt when non-empty"
  type        = string
  default     = ""
}
