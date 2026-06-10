//ff:func feature=cli type=command control=sequence
//ff:what 임시 블로그 루트(blog.yaml + 발행 글 1편)를 만들어 generate/check 테스트 픽스처로 제공
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeBlogFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [ko]\nsections: [opinion]\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	postDir := filepath.Join(dir, "content", "ko", "opinion")
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	post := "---\ntitle: Hello\ndate: 2026-01-02\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(postDir, "hello.md"), []byte(post), 0o644); err != nil {
		t.Fatalf("write post: %v", err)
	}
	return dir
}
