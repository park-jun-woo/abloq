//ff:func feature=quest type=parser control=sequence
//ff:what 파싱본(Doc) front matter에서 날짜 키 1개를 time.Time으로 해석 — fm-mirror(⑦)와 Seed lastmod 선별이 공유
package translation

import (
	"time"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// fmTime reads one front matter date key from a parsed Doc using the same
// interpreters the gate schema rule uses (gate.FMMap + gate.ParseFMTime).
func fmTime(d *agate.Doc, key string) (time.Time, bool) {
	if d == nil || !d.HasFM {
		return time.Time{}, false
	}
	m, ok := agate.FMMap(d.FrontMatter)
	if !ok {
		return time.Time{}, false
	}
	return agate.ParseFMTime(m[key])
}
