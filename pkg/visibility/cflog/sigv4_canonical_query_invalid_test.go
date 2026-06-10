//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what canonicalQuery가 파싱 불가 쿼리를 빈 정준 쿼리로 처리하는지 검증
package cflog

import "testing"

func TestCanonicalQueryInvalid(t *testing.T) {
	if got := canonicalQuery("k=%zz;bad"); got != "" {
		t.Errorf("canonicalQuery(invalid) = %q, want empty", got)
	}
}
