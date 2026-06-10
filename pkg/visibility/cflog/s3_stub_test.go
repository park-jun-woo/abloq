//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3 스텁 서버 — ListObjectsV2 2페이지 페이지네이션과 GetObject를 흉내, 서명 헤더 존재를 기록 (httptest)
package cflog

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newS3Stub serves a two-page ListObjectsV2 (logs/a.gz then logs/b.gz) and a
// gzip GetObject body, recording the last Authorization header so tests can
// assert the request was signed.
func newS3Stub(t *testing.T, lastAuth *string) *httptest.Server {
	t.Helper()
	var gzBody bytes.Buffer
	zw := gzip.NewWriter(&gzBody)
	zw.Write([]byte("2026-06-01\t12:00:01\tICN\t10\tip\tGET\thost\t/a/\t200\t-\tGPTBot/1.0\t-\t-\tHit\trid\n"))
	zw.Close()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*lastAuth = r.Header.Get("Authorization")
		q := r.URL.Query()
		if q.Get("list-type") == "2" {
			if q.Get("continuation-token") == "" {
				fmt.Fprint(w, `<?xml version="1.0"?><ListBucketResult><IsTruncated>true</IsTruncated><NextContinuationToken>tok2</NextContinuationToken><Contents><Key>logs/E.2026-06-01-12.a.gz</Key></Contents></ListBucketResult>`)
				return
			}
			fmt.Fprint(w, `<?xml version="1.0"?><ListBucketResult><IsTruncated>false</IsTruncated><Contents><Key>logs/E.2026-06-01-13.b.gz</Key></Contents></ListBucketResult>`)
			return
		}
		if r.URL.Path == "/logs/missing.gz" {
			http.Error(w, "<Error><Code>NoSuchKey</Code></Error>", http.StatusNotFound)
			return
		}
		w.Write(gzBody.Bytes())
	}))
}
