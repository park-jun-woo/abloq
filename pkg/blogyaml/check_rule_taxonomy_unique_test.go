//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleTaxonomyUnique 케이스 하나를 실행해 진단 수와 룰ID를 검증
package blogyaml

import "testing"

func checkRuleTaxonomyUnique(t *testing.T, taxonomy []string, wantDiags int) {
	t.Helper()
	b := &Blog{Geo: Geo{Taxonomy: taxonomy}}
	diags := ruleTaxonomyUnique("blog.yaml", b, lineIndex{})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "taxonomy-unique" {
			t.Errorf("want rule taxonomy-unique, got %q", d.Rule)
		}
	}
}
