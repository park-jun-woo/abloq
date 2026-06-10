//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what 로그 객체 키에서 UTC 시간 프리픽스(YYYY-MM-DD-HH) 추출 — 없으면 false (로그 파일이 아님)
package cflog

import "regexp"

var hourOfKeyRE = regexp.MustCompile(`\.(\d{4}-\d{2}-\d{2}-\d{2})\.`)

// hourOfKey extracts the UTC hour a CloudFront log object covers from its
// key (<prefix><dist-id>.YYYY-MM-DD-HH.<random>.gz). Keys without the hour
// component are not log objects and return false.
func hourOfKey(key string) (string, bool) {
	m := hourOfKeyRE.FindStringSubmatch(key)
	if m == nil {
		return "", false
	}
	return m[1], true
}
