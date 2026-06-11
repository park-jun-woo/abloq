//ff:func feature=quest type=parser control=sequence
//ff:what Prepare 검증 — 제출 JSON 디코드, 번역 단일 글 Target(Base nil) 조립, 원문 적재, 언어쌍 전달
package translation

import "testing"

func TestPrepare(t *testing.T) {
	orig := hugoLook
	hugoLook = func(string) (string, error) { return "/usr/bin/true", nil }
	defer func() { hugoLook = orig }()
	root := writeInstance(t)
	origin, ko := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	items, err := Definition{}.Seed([]string{originPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	writeFile(t, root, "content/ko/posts/fixture.md", ko)
	raw := []byte(`{"article":"content/ko/posts/fixture.md"}`)
	ctx, short, err := Definition{}.Prepare(nil, items[0], raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if short != nil {
		t.Fatalf("short verdict = %v, want nil", short)
	}
	sub := ctx.Submission.(*Submission)
	if len(sub.Target.Articles) != 1 || sub.Target.Articles[0].Base != nil {
		t.Errorf("target: %d article(s) — want 1 with Base nil", len(sub.Target.Articles))
	}
	if sub.Origin == nil || sub.Origin.Path != "content/en/posts/fixture.md" {
		t.Errorf("Origin = %+v, want the loaded en origin", sub.Origin)
	}
	if sub.Lang != "ko" || sub.OriginLang != "en" || sub.Root != root {
		t.Errorf("pair = %s->%s root=%s", sub.OriginLang, sub.Lang, sub.Root)
	}
	if ctx.Source == "" {
		t.Error("Source not filled with the translation body")
	}
}
