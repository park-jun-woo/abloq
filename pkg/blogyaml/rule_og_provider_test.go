//ff:func feature=blogyaml type=rule control=iteration dimension=1
//ff:what ruleOGProvider가 local/gemini/미선언을 통과시키고 미지 provider만 거부하는지 검증
package blogyaml

import (
	"strings"
	"testing"
)

func TestRuleOGProvider(t *testing.T) {
	for _, ok := range []string{"", "local", "gemini"} {
		b := &Blog{Image: Image{OG: ImageOG{Provider: ok}}}
		if diags := ruleOGProvider("blog.yaml", b, lineIndex{}); len(diags) != 0 {
			t.Errorf("provider %q: want 0 diagnostics, got %v", ok, diags)
		}
	}
	b := &Blog{Image: Image{OG: ImageOG{Provider: "dalle"}}}
	diags := ruleOGProvider("blog.yaml", b, lineIndex{"image.og.provider": 9})
	if len(diags) != 1 || diags[0].Rule != "og-provider" || diags[0].Line != 9 {
		t.Fatalf("provider dalle: %v", diags)
	}
	if !strings.Contains(diags[0].Message, "dalle") {
		t.Errorf("message = %q", diags[0].Message)
	}
}
