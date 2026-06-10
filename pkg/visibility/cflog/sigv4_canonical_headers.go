//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what 요청의 전체 헤더(+host)를 SigV4 정준 형식으로 — 소문자 이름 정렬, 값 트림, "k:v\n" 연접과 세미콜론 서명 목록
package cflog

import (
	"net/http"
	"sort"
	"strings"
)

// canonicalHeaders folds every header on the request plus host into the
// SigV4 canonical form. It returns the canonical block (terminated by \n)
// and the signed-headers list.
func canonicalHeaders(req *http.Request) (string, string) {
	host := req.Host
	if host == "" {
		host = req.URL.Host
	}
	byName := map[string]string{"host": strings.TrimSpace(host)}
	for name, vals := range req.Header {
		byName[strings.ToLower(name)] = strings.TrimSpace(strings.Join(vals, ","))
	}
	names := make([]string, 0, len(byName))
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)
	var block, signed strings.Builder
	for i, name := range names {
		block.WriteString(name + ":" + byName[name] + "\n")
		if i > 0 {
			signed.WriteString(";")
		}
		signed.WriteString(name)
	}
	return block.String(), signed.String()
}
