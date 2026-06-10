//ff:func feature=gen type=generator control=sequence
//ff:what Write가 디렉토리 생성 실패(경로가 파일)와 파일 기록 실패(경로가 디렉토리)에서 에러를 반환하는지 검증
package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteError(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "static"), []byte("file in the way"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := Write(dir, []Output{{Path: "static/robots.txt", Data: []byte("x")}}); err == nil {
		t.Errorf("want MkdirAll error when 'static' is a file, got nil")
	}
	if err := os.MkdirAll(filepath.Join(dir, "hugo.toml"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := Write(dir, []Output{{Path: "hugo.toml", Data: []byte("x")}}); err == nil {
		t.Errorf("want WriteFile error when 'hugo.toml' is a directory, got nil")
	}
}
