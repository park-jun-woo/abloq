//ff:func feature=quest type=generator control=sequence
//ff:what Render 검증 — 프롬프트에 대상 경로·언어쌍·헤딩 맵(원문→대상)·원문 전문·tasks/context 문서가 포함되는지
package translation

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	root := writeInstance(t)
	origin, _ := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	items, err := Definition{}.Seed([]string{originPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	out, err := Definition{}.Render(nil, items[0])
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	for _, want := range []string{
		"translation quest — ko/posts/fixture",
		"target article (ko): content/ko/posts/fixture.md",
		`- sources: "Sources" -> "출처"`,
		"Intro paragraph with an [external spec]",
		"번역 퀘스트 — 태스크 트리",
		"번역 퀘스트 — 컨텍스트 규약",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
}
