//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what 쿼리 문자열을 SigV4 정준 형식으로 — 키·값 RFC3986 인코딩, 키(동일 키는 값) 사전순 정렬, '&' 연접
package cflog

import (
	"net/url"
	"sort"
	"strings"
)

// canonicalQuery rewrites a raw query string into the SigV4 canonical form:
// every key and value percent-encoded with unreserved characters bare,
// pairs sorted by encoded key then value.
func canonicalQuery(rawQuery string) string {
	vals, err := url.ParseQuery(rawQuery)
	if err != nil {
		vals = url.Values{}
	}
	var pairs []string
	for key, vv := range vals {
		for _, v := range vv {
			pairs = append(pairs, awsEscape(key, false)+"="+awsEscape(v, false))
		}
	}
	sort.Strings(pairs)
	return strings.Join(pairs, "&")
}
