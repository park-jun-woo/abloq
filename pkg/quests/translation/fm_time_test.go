//ff:func feature=quest type=parser control=sequence
//ff:what fmTime 검증 — date/lastmod 해석, 키 부재·front matter 없음·nil Doc은 false
package translation

import (
	"testing"
	"time"
)

func TestFmTime(t *testing.T) {
	origin, _ := passPair()
	d := docOf(t, "en", origin)
	got, ok := fmTime(d, "lastmod")
	if !ok || !got.Equal(time.Date(2026, 6, 3, 0, 0, 0, 0, time.UTC)) {
		t.Errorf("lastmod = %v ok=%v", got, ok)
	}
	if _, ok := fmTime(d, "nope"); ok {
		t.Error("absent key: want ok=false")
	}
	if _, ok := fmTime(docOf(t, "en", "no front matter\n"), "date"); ok {
		t.Error("no front matter: want ok=false")
	}
	if _, ok := fmTime(nil, "date"); ok {
		t.Error("nil doc: want ok=false")
	}
	broken := docOf(t, "en", "---\ntitle: [unclosed\n---\n\nbody\n")
	if _, ok := fmTime(broken, "date"); ok {
		t.Error("broken YAML front matter: want ok=false")
	}
}
