//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3Source.Get이 Prefix 붙인 키의 본문을 스트림하고 NoSuchKey는 에러인지 검증 (스텁)
package cflog

import (
	"testing"
)

func TestS3Get(t *testing.T) {
	var lastAuth string
	stub := newS3Stub(t, &lastAuth)
	defer stub.Close()
	src := S3Source{
		Bucket: "b", Prefix: "logs/", Region: "us-east-1",
		AccessKey: "AKIDEXAMPLE", SecretKey: "secret",
		Endpoint: stub.URL, Client: stub.Client(),
	}
	rc, err := src.Get("E.2026-06-01-12.a.gz")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	recs, err := parseRecords(rc)
	rc.Close()
	if err != nil || len(recs) != 1 || recs[0].URI != "/a/" {
		t.Errorf("recs = %+v, err %v", recs, err)
	}
	if _, err := src.Get("missing.gz"); err == nil {
		t.Errorf("NoSuchKey accepted")
	}
}
