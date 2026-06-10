//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 글 집계 대상 상태 코드 판정 — 2xx 또는 304만 true
package cflog

import "strings"

// statusOK reports whether a hit with this sc-status counts toward article
// aggregation: successful responses (2xx) and not-modified revalidations
// (304).
func statusOK(status string) bool {
	return (len(status) == 3 && strings.HasPrefix(status, "2")) || status == "304"
}
