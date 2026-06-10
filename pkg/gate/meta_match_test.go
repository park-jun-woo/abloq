//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what metaMatch 케이스 — title 일치, og:title 일치, 역방향(짧은 title ⊂ 긴 표기) 일치, 불일치, 토큰 없는 표기는 통과
package gate

import "testing"

func TestMetaMatch(t *testing.T) {
	cases := []struct {
		name, html, label string
		want              bool
	}{
		{"title matches", `<title>Example Benchmark Report</title>`, "Example Benchmark Report", true},
		{"og title matches", `<meta property="og:title" content="Example Benchmark Report">`, "example benchmark", true},
		{"short title inside long label", `<title>GEO</title>`, "GEO paper by Aggarwal and others twenty twenty four", true},
		{"mismatch", `<title>전혀 다른 페이지</title>`, "Example Benchmark Report", false},
		{"no titles at all", `<p>no head</p>`, "Example Benchmark Report", false},
		{"bare url label passes", `<title>Whatever</title>`, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := metaMatch(tc.html, tc.label); got != tc.want {
				t.Errorf("metaMatch(label %q) = %v, want %v", tc.label, got, tc.want)
			}
		})
	}
}
