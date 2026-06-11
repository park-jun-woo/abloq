//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what llmsTxtItemLine이 항목 경로의 라인을 우선 반환하고, 없으면 상위 키 → geo.llms_txt 순으로 폴백하는지 검증
package blogyaml

import "testing"

func TestLlmsTxtItemLine(t *testing.T) {
	idx := lineIndex{
		"geo.llms_txt":               4,
		"geo.llms_txt.pinned":        5,
		"geo.llms_txt.pinned[0].url": 7,
	}
	cases := []struct {
		name, item, parent string
		want               int
	}{
		{"item present", "pinned[0].url", "pinned", 7},
		{"absent item falls back to parent", "pinned[1].url", "pinned", 5},
		{"absent parent falls back to llms_txt", "languages[0]", "languages", 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := llmsTxtItemLine(idx, tc.item, tc.parent); got != tc.want {
				t.Errorf("llmsTxtItemLine(%q, %q) = %d, want %d", tc.item, tc.parent, got, tc.want)
			}
		})
	}
}
