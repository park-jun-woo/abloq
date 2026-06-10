//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what s3Do가 비-2xx 응답을 본문 머리를 담은 에러로 바꾸는지, baseURL이 Endpoint 오버라이드와 가상 호스트 기본값을 푸는지 검증
package cflog

import (
	"strings"
	"testing"
)

func TestS3Do(t *testing.T) {
	var lastAuth string
	stub := newS3Stub(t, &lastAuth)
	defer stub.Close()
	src := S3Source{
		Bucket: "b", Prefix: "logs/", Region: "ap-northeast-2",
		AccessKey: "AKIDEXAMPLE", SecretKey: "secret",
		Endpoint: stub.URL, Client: stub.Client(),
	}
	if _, err := src.s3Do(stub.URL + "/logs/missing.gz"); err == nil || !strings.Contains(err.Error(), "NoSuchKey") {
		t.Errorf("err = %v, want NoSuchKey body head", err)
	}
	if got := src.baseURL(); got != stub.URL {
		t.Errorf("baseURL = %q, want endpoint override", got)
	}
	src.Endpoint = ""
	if got := src.baseURL(); got != "https://b.s3.ap-northeast-2.amazonaws.com" {
		t.Errorf("baseURL = %q", got)
	}
}
