//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what front matter 날짜 키의 실값 표시 문자열 — 파싱되면 ISO 표기, 아니면 "missing or unparseable" (fm-mirror Fact의 Actual)
package translation

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// fmValue renders one front matter date key for Fact display.
func fmValue(d *agate.Doc, key string) string {
	t, ok := fmTime(d, key)
	if !ok {
		return key + " missing or unparseable"
	}
	return t.Format("2006-01-02T15:04:05Z07:00")
}
