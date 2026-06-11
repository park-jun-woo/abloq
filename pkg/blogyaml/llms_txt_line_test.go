//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what llmsTxtLine이 하위 키 라인을 우선 반환하고, 없으면 geo.llms_txt 키 라인 → 1 순으로 폴백하는지 검증
package blogyaml

import "testing"

func TestLlmsTxtLine(t *testing.T) {
	cases := []struct {
		name string
		idx  lineIndex
		key  string
		want int
	}{
		{"sub-key present", lineIndex{"geo.llms_txt": 4, "geo.llms_txt.mode": 5}, "mode", 5},
		{"absent sub-key falls back to llms_txt", lineIndex{"geo.llms_txt": 4}, "mode", 4},
		{"shorthand without llms_txt key falls back to 1", lineIndex{}, "mode", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := llmsTxtLine(tc.idx, tc.key); got != tc.want {
				t.Errorf("llmsTxtLine(%q) = %d, want %d", tc.key, got, tc.want)
			}
		})
	}
}
