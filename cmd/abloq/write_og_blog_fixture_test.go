//ff:func feature=cli type=command control=sequence
//ff:what 테스트 픽스처 — image.og(gemini + 안 3종: minimal/photo/bad)를 선언한 유효한 blog.yaml 임시 블로그 루트 생성
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeOGBlogFixture(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	blogYAML := "site:\n  baseURL: https://t.example.com\n  title: T\n  author: A\n" +
		"languages: [en]\nsections: [tech]\n" +
		"image:\n  og:\n    provider: gemini\n    model: good-model\n" +
		"    prompt: Site prompt for {title}\n" +
		"    variants:\n" +
		"      - name: minimal\n" +
		"      - name: photo\n        model: photo-model\n        overlay: true\n" +
		"      - name: bad\n        model: boom-model\n"
	if err := os.WriteFile(filepath.Join(dir, "blog.yaml"), []byte(blogYAML), 0o644); err != nil {
		t.Fatalf("write blog.yaml: %v", err)
	}
	return dir
}
