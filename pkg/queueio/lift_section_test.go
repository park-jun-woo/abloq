//ff:func feature=queueio type=parser control=sequence
//ff:what liftSection이 section을 분리하고 원본 맵을 변형하지 않는지 검증
package queueio

import "testing"

func TestLiftSection(t *testing.T) {
	in := map[string]string{"section": "tech", "lastmod": "2026-06-05"}
	section, rest := liftSection(in)
	if section != "tech" {
		t.Errorf("want tech, got %s", section)
	}
	if _, dup := rest["section"]; dup || rest["lastmod"] != "2026-06-05" {
		t.Errorf("unexpected rest: %+v", rest)
	}
	if in["section"] != "tech" {
		t.Error("input map must stay untouched")
	}
}
