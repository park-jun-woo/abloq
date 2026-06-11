//ff:func feature=cli type=command control=sequence
//ff:what 테스트 픽스처 — base 언어(en) front matter slug 오버라이드 글 + 같은 slug의 ko 번역(base만 매칭됨을 검증)을 둔 임시 블로그 루트 생성
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSummaryFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [en, ko]\nsections: [tech]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	enDir := filepath.Join(dir, "content", "en", "tech")
	if err := os.MkdirAll(enDir, 0o755); err != nil {
		t.Fatalf("MkdirAll en: %v", err)
	}
	koDir := filepath.Join(dir, "content", "ko", "tech")
	if err := os.MkdirAll(koDir, 0o755); err != nil {
		t.Fatalf("MkdirAll ko: %v", err)
	}
	// base-language article whose front matter slug overrides the file stem
	overridden := "---\ntitle: Overridden\ndate: 2026-01-02\nslug: real-slug\nsummary: base summary text\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(enDir, "file-stem.md"), []byte(overridden), 0o644); err != nil {
		t.Fatalf("write base post: %v", err)
	}
	// a translation in ko reuses the same slug — must NOT be matched (base only)
	ko := "---\ntitle: 번역\ndate: 2026-01-02\nslug: real-slug\nsummary: 번역 요약\n---\n본문\n"
	if err := os.WriteFile(filepath.Join(koDir, "real-slug.md"), []byte(ko), 0o644); err != nil {
		t.Fatalf("write ko post: %v", err)
	}
	return dir
}
