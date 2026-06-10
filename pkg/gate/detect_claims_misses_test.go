//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 알려진 미탐 목록(claims-misses.md) 고정 — 수사적 수치/단정 없는 표기/통화 접두/단위 없는 수는 현재 검출하지 않음
//ff:why 오탐 0이 미탐 0보다 우선 — 미탐 패턴은 픽스처로 목록화해 검출기 강화 시 회귀 기준으로 쓴다
package gate

import "testing"

func TestDetectClaimsMisses(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "misses", "evidence/claims-misses.md")
	if got := DetectClaims(a.Doc); len(got) != 0 {
		t.Errorf("claims-misses.md documents known false negatives; detector now finds %d — move them to claims-corpus.md: %+v", len(got), got)
	}
}
