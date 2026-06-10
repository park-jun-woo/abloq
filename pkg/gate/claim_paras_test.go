//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what claimParas 케이스 — 빈 줄 문단 구분, 코드 펜스/들여쓴 코드/인용/헤딩/이미지/구조 라인 제외 검증
package gate

import "testing"

func TestClaimParas(t *testing.T) {
	b := loadGateBlog(t)
	content := "![Main](/i.webp)\n*Image: x*\n\nFirst line.\nSecond line.\n\n" +
		"```\nfenced 42% improved\n```\n\n> quoted 42% improved\n\n    indented 42% improved\n\n" +
		"## Sources\n\n- item\n\nLast para.\n"
	paras := claimParas(ParseArticle(b, "en", content))
	if len(paras) != 3 {
		t.Fatalf("want 3 paragraphs (code/quote/heading/image excluded), got %d: %+v", len(paras), paras)
	}
	cases := []struct {
		name      string
		idx       int
		wantLen   int
		wantFirst string
	}{
		{"two-line paragraph kept together", 0, 2, "First line."},
		{"sources list item is prose for the detector", 1, 1, "- item"},
		{"trailing paragraph", 2, 1, "Last para."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := paras[tc.idx]
			if len(p.texts) != tc.wantLen || p.texts[0] != tc.wantFirst {
				t.Errorf("paras[%d] = %v, want %d line(s) starting %q", tc.idx, p.texts, tc.wantLen, tc.wantFirst)
			}
		})
	}
}
