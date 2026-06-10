//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what s3Do가 잘못된 URL과 연결 불가 endpoint를 에러로 전파하는지 검증
package cflog

import "testing"

func TestS3DoTransportErrors(t *testing.T) {
	src := S3Source{Bucket: "b", Region: "us-east-1", AccessKey: "k", SecretKey: "s"}
	if _, err := src.s3Do("http://127.0.0.1:1/\x7f"); err == nil {
		t.Error("invalid URL accepted")
	}
	if _, err := src.s3Do("http://127.0.0.1:1/unreachable"); err == nil {
		t.Error("unreachable endpoint accepted")
	}
}
