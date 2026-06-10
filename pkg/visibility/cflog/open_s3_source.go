//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what s3://bucket/prefix 명세를 env 자격증명(AWS_ACCESS_KEY_ID/SECRET/REGION/SESSION_TOKEN)의 S3 소스로 — 자격 결손은 에러, CF_LOG_S3_ENDPOINT는 로컬 스텁 오버라이드
package cflog

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// openS3Source builds an S3Source from "s3://bucket/prefix" (rest is the
// spec without the scheme) and the standard AWS env credentials.
func openS3Source(spec, rest string) (Source, error) {
	bucket, prefix, _ := strings.Cut(rest, "/")
	if bucket == "" {
		return nil, fmt.Errorf("cflog: no bucket in source %q", spec)
	}
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
	access := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if region == "" || access == "" || secret == "" {
		return nil, errors.New("cflog: s3 source needs AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY and AWS_REGION")
	}
	return S3Source{
		Bucket:       bucket,
		Prefix:       prefix,
		Region:       region,
		AccessKey:    access,
		SecretKey:    secret,
		SessionToken: os.Getenv("AWS_SESSION_TOKEN"),
		Endpoint:     os.Getenv("CF_LOG_S3_ENDPOINT"),
	}, nil
}
