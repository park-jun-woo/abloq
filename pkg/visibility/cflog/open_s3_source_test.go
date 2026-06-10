//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what openS3Source가 버킷/프리픽스와 env 자격증명을 S3Source로 풀고, 버킷·자격 결손과 AWS_DEFAULT_REGION 폴백을 처리하는지 검증
package cflog

import "testing"

func TestOpenS3Source(t *testing.T) {
	if _, err := openS3Source("s3://", ""); err == nil {
		t.Error("bucketless spec accepted")
	}
	t.Setenv("AWS_ACCESS_KEY_ID", "AKIDEXAMPLE")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	t.Setenv("AWS_REGION", "")
	t.Setenv("AWS_DEFAULT_REGION", "")
	if _, err := openS3Source("s3://b/logs/", "b/logs/"); err == nil {
		t.Error("regionless credentials accepted")
	}
	t.Setenv("AWS_DEFAULT_REGION", "us-east-1")
	src, err := openS3Source("s3://b/logs/", "b/logs/")
	if err != nil {
		t.Fatalf("openS3Source: %v", err)
	}
	s3 := src.(S3Source)
	if s3.Bucket != "b" || s3.Prefix != "logs/" || s3.Region != "us-east-1" {
		t.Errorf("s3 = %+v (AWS_DEFAULT_REGION fallback expected)", s3)
	}
}
