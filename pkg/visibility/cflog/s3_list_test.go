//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3Source.List가 페이지네이션을 따라가 Prefix를 벗긴 정렬 키를 모으고 서명 헤더를 붙이는지 검증 (스텁)
package cflog

import (
	"reflect"
	"strings"
	"testing"
)

func TestS3List(t *testing.T) {
	var lastAuth string
	stub := newS3Stub(t, &lastAuth)
	defer stub.Close()
	src := S3Source{
		Bucket: "b", Prefix: "logs/", Region: "us-east-1",
		AccessKey: "AKIDEXAMPLE", SecretKey: "secret",
		Endpoint: stub.URL, Client: stub.Client(),
	}
	got, err := src.List("", "")
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	want := []string{"E.2026-06-01-12.a.gz", "E.2026-06-01-13.b.gz"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("List = %v, want %v", got, want)
	}
	if !strings.HasPrefix(lastAuth, "AWS4-HMAC-SHA256 Credential=AKIDEXAMPLE/") {
		t.Errorf("request not signed: Authorization = %q", lastAuth)
	}
}
