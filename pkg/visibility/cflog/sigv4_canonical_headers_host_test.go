//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what canonicalHeaders가 명시 req.Host를 URL 호스트보다 우선하는지 검증
package cflog

import (
	"net/http"
	"testing"
)

func TestCanonicalHeadersExplicitHost(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://b.s3.amazonaws.com/", nil)
	req.Host = "override.example.com"
	block, _ := canonicalHeaders(req)
	if block != "host:override.example.com\n" {
		t.Errorf("block = %q, want the explicit req.Host", block)
	}
}
