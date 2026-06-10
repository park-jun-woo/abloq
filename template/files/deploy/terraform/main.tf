# abloq blog infrastructure — S3 + CloudFront (+ optional WAF, CF logs, IndexNow key)

terraform {
  required_version = ">= 1.5.0"

  required_providers {
    aws = { source = "hashicorp/aws", version = "~> 5.0" }
  }
}

provider "aws" {
  region = var.aws_region
}

# CloudFront-scoped resources (ACM cert, WAF) must live in us-east-1.
provider "aws" {
  alias  = "us_east_1"
  region = "us-east-1"
}

module "logs" {
  count  = var.enable_logs ? 1 : 0
  source = "./modules/logs"

  project        = var.project
  retention_days = var.log_retention_days
}

module "waf" {
  count  = var.enable_waf ? 1 : 0
  source = "./modules/waf"

  providers = { aws = aws.us_east_1 }

  project    = var.project
  rate_limit = var.waf_rate_limit
}

module "site" {
  source = "./modules/site"

  providers = {
    aws           = aws
    aws.us_east_1 = aws.us_east_1
  }

  project            = var.project
  domain             = var.domain
  web_acl_arn        = var.enable_waf ? module.waf[0].web_acl_arn : ""
  logs_bucket_domain = var.enable_logs ? module.logs[0].bucket_domain_name : ""
  indexnow_key       = var.indexnow_key
}
