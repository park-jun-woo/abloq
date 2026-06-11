//ff:func feature=quest type=rule control=sequence
//ff:what hasDisposition 검증 — addressed/revised/excluded 토큰만 인정, 다른 토큰·접두 불일치는 불인정
package writing

import "testing"

func TestHasDisposition(t *testing.T) {
	if !hasDisposition("- c1: addressed — ok\n", "c1") {
		t.Error("addressed: want true")
	}
	if !hasDisposition("  - c1: excluded out of scope\n", "c1") {
		t.Error("indented excluded: want true")
	}
	if hasDisposition("- c1: maybe — unsure\n", "c1") {
		t.Error("unknown token: want false")
	}
	if hasDisposition("- c10: addressed\n", "c1") {
		t.Error("id prefix mismatch: want false")
	}
	if hasDisposition("- c1:\n", "c1") {
		t.Error("empty disposition: want false")
	}
}
