variable "project" {
  description = "Project name used as resource prefix"
  type        = string
}

variable "rate_limit" {
  description = "Requests per 5 minutes per IP before blocking"
  type        = number
  default     = 2000
}
