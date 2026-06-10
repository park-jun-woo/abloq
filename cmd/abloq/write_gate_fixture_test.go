//ff:func feature=cli type=command control=iteration dimension=1
//ff:what 임시 게이트 픽스처 저장소 생성 — 정규(전 룰 통과) 또는 위반(이미지 누락) 글 1쌍(ko/en)
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeGateFixture(t *testing.T, canonical bool) string {
	t.Helper()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko, en]\nsections: [tech]\n" +
		"structure:\n  order: [image, attribution, body, sources]\n" +
		"  headings:\n    sources: { ko: \"출처\", en: \"Sources\" }\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	body := "본문 텍스트.\n"
	if canonical {
		body = "![Hello](/i.webp)\n*Image: AI generated*\n\n본문 텍스트.\n"
	}
	post := "---\ntitle: Hello\ndate: 2026-01-02\nlastmod: 2026-01-03\ntags: [\"a\"]\n---\n\n" + body
	for _, lang := range []string{"ko", "en"} {
		postDir := filepath.Join(dir, "content", lang, "tech")
		if err := os.MkdirAll(postDir, 0o755); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(filepath.Join(postDir, "hello.md"), []byte(post), 0o644); err != nil {
			t.Fatalf("write post: %v", err)
		}
	}
	return dir
}
