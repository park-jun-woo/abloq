//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what paraClaims 케이스 — 출처 링크가 있는 문단의 주장은 Sourced, 없는 문단은 unsourced, 파일 라인 번호 검증
package gate

import "testing"

func TestParaClaims(t *testing.T) {
	b := loadGateBlog(t)
	d := ParseArticle(b, "en", "Throughput improved by 42%.\nSee the [bench](https://example.com/b).\n\nLatency dropped to 12ms.\n")
	paras := claimParas(d)
	if len(paras) != 2 {
		t.Fatalf("want 2 paragraphs, got %d", len(paras))
	}
	sourced := paraClaims(d, paras[0])
	if len(sourced) != 1 || !sourced[0].Sourced || sourced[0].Line != 1 {
		t.Errorf("sourced paragraph: got %+v, want 1 sourced claim at line 1", sourced)
	}
	bare := paraClaims(d, paras[1])
	if len(bare) != 1 || bare[0].Sourced || bare[0].Text != "Latency dropped to 12ms." {
		t.Errorf("bare paragraph: got %+v, want 1 unsourced claim", bare)
	}
}
