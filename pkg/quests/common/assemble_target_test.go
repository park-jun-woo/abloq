//ff:func feature=quest type=frame control=sequence
//ff:what AssembleTarget 검증 — blog.yaml+글 1편으로 Target(Base nil) 조립과 원문 바이트 반환, 글 부재는 에러
package common

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixtureArticleMD = `---
title: "A"
date: 2026-06-01
lastmod: 2026-06-02
tags: [t]
---

Body text.

## Sources

- [Spec](https://example.org/spec)
`

func TestAssembleTarget(t *testing.T) {
	root, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	tgt, body, err := AssembleTarget(root, "content/en/posts/a.md", "en", "posts", "a")
	if err != nil {
		t.Fatalf("AssembleTarget: %v", err)
	}
	if len(tgt.Articles) != 1 || tgt.Articles[0].Base != nil {
		t.Fatalf("articles = %d — want 1 with Base nil", len(tgt.Articles))
	}
	if tgt.Articles[0].Slug != "a" || len(body) == 0 {
		t.Errorf("slug = %q, body %d byte(s)", tgt.Articles[0].Slug, len(body))
	}
	if _, _, err := AssembleTarget(root, "content/en/posts/missing.md", "en", "posts", "missing"); err == nil {
		t.Error("missing article: want error")
	}
	if _, _, err := AssembleTarget(t.TempDir(), "content/en/posts/a.md", "en", "posts", "a"); err == nil {
		t.Error("no blog.yaml: want load error")
	}
	bad, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	if err := os.WriteFile(filepath.Join(bad, "blog.yaml"), []byte(strings.Replace(fixtureBlogYAML, "https://example.com", "not-a-url", 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := AssembleTarget(bad, "content/en/posts/a.md", "en", "posts", "a"); err == nil {
		t.Error("blog.yaml with diagnostics: want error")
	}
}
