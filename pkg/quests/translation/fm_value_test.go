//ff:func feature=quest type=rule control=sequence
//ff:what fmValue 검증 — 파싱 키는 ISO 표기, 부재 키는 "missing or unparseable" 문구
package translation

import (
	"strings"
	"testing"
)

func TestFmValue(t *testing.T) {
	origin, _ := passPair()
	d := docOf(t, "en", origin)
	if got := fmValue(d, "lastmod"); !strings.HasPrefix(got, "2026-06-03") {
		t.Errorf("lastmod = %q", got)
	}
	if got := fmValue(d, "nope"); !strings.Contains(got, "missing or unparseable") {
		t.Errorf("absent = %q", got)
	}
}
