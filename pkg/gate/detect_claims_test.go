//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what DetectClaims 코퍼스 검증 — 양성 5건만 정확히 검출(오탐 0), 전부 unsourced, 음성(연도·버전·코드·인용)은 0건
package gate

import "testing"

func TestDetectClaimsCorpus(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "corpus", "evidence/claims-corpus.md")
	want := []string{
		"Cache hit rate improved by 38% after the rollout.",
		"응답 시간이 120ms로 줄었다.",
		"전환율이 3배 증가했다.",
		"Memory usage dropped to 512 MB under the new allocator.",
		"시장 점유율 45%를 기록했다.",
	}
	claims := DetectClaims(a.Doc)
	if len(claims) != len(want) {
		t.Fatalf("want exactly %d claims (zero false positives), got %d: %+v", len(want), len(claims), claims)
	}
	for i, c := range claims {
		if c.Text != want[i] {
			t.Errorf("claims[%d].Text = %q, want %q", i, c.Text, want[i])
		}
		if c.Sourced {
			t.Errorf("corpus has no source links — claim %q must be unsourced", c.Text)
		}
	}
}
