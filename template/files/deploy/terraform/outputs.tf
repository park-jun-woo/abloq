output "bucket_name" {
  description = "S3 origin bucket — sync public/ here"
  value       = module.site.bucket_name
}

output "distribution_id" {
  description = "CloudFront distribution id — invalidate after deploy"
  value       = module.site.distribution_id
}

output "distribution_domain" {
  description = "CloudFront domain name"
  value       = module.site.distribution_domain
}

output "name_servers" {
  description = "Route53 zone name servers — set these at your registrar"
  value       = module.site.name_servers
}

output "logs_bucket" {
  description = "CloudFront log bucket name (empty when logs disabled)"
  value       = var.enable_logs ? module.logs[0].bucket_name : ""
}
