//ff:func feature=init type=generator control=sequence
//ff:what CopyFile이 상위 디렉토리를 만들며 파일을 기록하고 없는 소스에 에러를 내는지 검증
package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopyFile(t *testing.T) {
	fsys := fstest.MapFS{"deep/file.txt": {Data: []byte("hello")}}
	dir := t.TempDir()
	dst := filepath.Join(dir, "deep", "file.txt")
	if err := CopyFile(fsys, "deep/file.txt", dst); err != nil {
		t.Fatalf("CopyFile: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil || string(got) != "hello" {
		t.Errorf("read back = %q, err %v", got, err)
	}
	if err := CopyFile(fsys, "missing.txt", filepath.Join(dir, "x")); err == nil {
		t.Error("CopyFile(missing) expected error, got nil")
	}
	if err := CopyFile(fsys, "deep/file.txt", filepath.Join(dst, "under-file", "y")); err == nil {
		t.Error("CopyFile under a regular file expected mkdir error, got nil")
	}
}
