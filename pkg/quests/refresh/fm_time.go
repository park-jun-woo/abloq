//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 파싱본(Doc) front matter에서 날짜 키 1개를 time.Time으로 해석 — lastmod-advance 룰이 기준선·제출본 비교에 사용
package refresh

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
