//ff:func feature=quest type=generator control=iteration dimension=1
//ff:what Render 검증 — 프롬프트에 Key·대상 경로·insight 원문·tasks.md(T1~T4)·context.md(금지 사항)가 결합되는지
package writing

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	root := writeInstance(t)
	_, ins := passFixtures()
	path := writeFile(t, root, "content/en/posts/hello.insight.yaml", ins)
	items, err := Definition{}.Seed([]string{path})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	got, err := Definition{}.Render(nil, items[0])
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	for _, want := range []string{
		"writing quest — en/posts/hello",
		"target article: content/en/posts/hello.md",
		"submit --key en/posts/hello",
		"topic: \"test topic\"", // insight raw
		"## T1 — 자료 수집",         // tasks.md
		"## T4 — REVIEW",        // tasks.md
		"자가 REVIEW 금지",          // context.md
	} {
		if !strings.Contains(got, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
}
