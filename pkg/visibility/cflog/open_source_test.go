//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what OpenSource가 디렉토리는 DirSource로, s3://는 env 자격증명의 S3Source로 풀고 자격 결손·비디렉토리는 에러인지 검증
package cflog

import (
	"testing"
)

func TestOpenSource(t *testing.T) {
	dir := t.TempDir()
	src, err := OpenSource(dir)
	if err != nil {
		t.Fatalf("OpenSource(dir): %v", err)
	}
	if _, ok := src.(DirSource); !ok {
		t.Errorf("src = %T, want DirSource", src)
	}
	if _, err := OpenSource(dir + "/missing"); err == nil {
		t.Errorf("missing directory accepted")
	}

	t.Setenv("AWS_ACCESS_KEY_ID", "")
	if _, err := OpenSource("s3://bucket/logs/"); err == nil {
		t.Errorf("s3 source without credentials accepted")
	}
	t.Setenv("AWS_ACCESS_KEY_ID", "AKIDEXAMPLE")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	t.Setenv("AWS_REGION", "ap-northeast-2")
	t.Setenv("AWS_SESSION_TOKEN", "tok")
	src, err = OpenSource("s3://bucket/logs/")
	if err != nil {
		t.Fatalf("OpenSource(s3): %v", err)
	}
	s3, ok := src.(S3Source)
	if !ok {
		t.Fatalf("src = %T, want S3Source", src)
	}
	if s3.Bucket != "bucket" || s3.Prefix != "logs/" || s3.Region != "ap-northeast-2" || s3.SessionToken != "tok" {
		t.Errorf("s3 = %+v", s3)
	}
	if _, err := OpenSource("s3://"); err == nil {
		t.Errorf("bucketless s3 spec accepted")
	}
}
