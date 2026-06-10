# deploy/terraform — S3 + CloudFront static hosting (optional)

Optional IaC for the blog. Modules are independent toggles:

| module | toggle | what it creates |
|---|---|---|
| `modules/site` | always | private S3 bucket, CloudFront (OAC) distribution, ACM cert, Route53 zone + alias, IndexNow key file |
| `modules/waf` | `enable_waf` | WAFv2 web ACL (AWS managed common rules + rate limit), attached to CloudFront |
| `modules/logs` | `enable_logs` | CloudFront standard-log S3 bucket (input for the abloq visibility pipeline) |

## Usage

```bash
cd deploy/terraform
cp example.tfvars prod.tfvars   # edit values
terraform init
terraform plan  -var-file=prod.tfvars
terraform apply -var-file=prod.tfvars
```

Deploy the built site:

```bash
hugo && abloq gate . && abloq postbuild md
aws s3 sync public/ "s3://$(terraform output -raw bucket_name)/" --delete
aws cloudfront create-invalidation --distribution-id "$(terraform output -raw distribution_id)" --paths '/*'
```

`indexnow_key` must match `deploy.indexnow` usage: the key file is served at
`https://{domain}/{key}.txt` as IndexNow requires.
