//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleCrawlersPolicy 케이스 하나를 실행해 진단 수·룰ID·첫 진단 메시지를 검증
package blogyaml

import (
	"strings"
	"testing"
)

func checkRuleCrawlersPolicy(t *testing.T, crawlers map[string]string, wantDiags int, wantFirst string) {
	t.Helper()
	b := &Blog{Geo: Geo{Crawlers: crawlers}}
	diags := ruleCrawlersPolicy("blog.yaml", b, lineIndex{})
	if len(diags) != wantDiags {
		t.Fatalf("want %d diagnostics, got %d: %v", wantDiags, len(diags), diags)
	}
	for _, d := range diags {
		if d.Rule != "crawlers-policy" {
			t.Errorf("want rule crawlers-policy, got %q", d.Rule)
		}
	}
	if wantFirst != "" && !strings.Contains(diags[0].Message, wantFirst) {
		t.Errorf("want first diagnostic about %q, got %q", wantFirst, diags[0].Message)
	}
}
