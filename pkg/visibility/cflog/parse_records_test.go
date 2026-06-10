//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what parseRecords가 gzip 스트림에서 유효 행만 Record로 모으고 비-gzip 입력은 에러인지 검증
package cflog

import (
	"bytes"
	"compress/gzip"
	"strings"
	"testing"
)

func TestParseRecords(t *testing.T) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Write([]byte("#Fields: ...\n" +
		"2026-06-01\t12:00:01\tICN\t10\tip\tGET\thost\t/a/\t200\t-\tua\t-\t-\tHit\trid\n" +
		"bad line\n" +
		"2026-06-01\t12:00:02\tICN\t10\tip\tGET\thost\t/b/\t304\t-\tua\t-\t-\tHit\trid\n"))
	zw.Close()
	recs, err := parseRecords(&buf)
	if err != nil {
		t.Fatalf("parseRecords: %v", err)
	}
	if len(recs) != 2 || recs[0].URI != "/a/" || recs[1].Status != "304" {
		t.Errorf("recs = %+v", recs)
	}
	if _, err := parseRecords(strings.NewReader("not gzip")); err == nil {
		t.Errorf("non-gzip input accepted")
	}
}
