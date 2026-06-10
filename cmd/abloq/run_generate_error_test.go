//ff:func feature=cli type=command control=sequence
//ff:what runGenerate가 blog.yaml 누락과 파생물 기록 실패(경로 충돌)에서 에러를 반환하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunGenerateError(t *testing.T) {
	var out bytes.Buffer
	if err := runGenerate(&out, t.TempDir()); err == nil {
		t.Errorf("want error for dir without blog.yaml, got nil")
	}
	dir := writeBlogFixture(t)
	if err := os.WriteFile(filepath.Join(dir, "static"), []byte("file in the way"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := runGenerate(&out, dir); err == nil {
		t.Errorf("want write error when 'static' is a file, got nil")
	}
}
