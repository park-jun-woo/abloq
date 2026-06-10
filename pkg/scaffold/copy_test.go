//ff:func feature=init type=generator control=iteration dimension=1
//ff:what Copy가 fs 전체를 상대 경로 보존 복제하고 파일 수를 반환하는지, 재실행이 멱등인지 검증
package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopy(t *testing.T) {
	fsys := fstest.MapFS{
		"README.md":  {Data: []byte("hi")},
		"a/b/c.html": {Data: []byte("<x>")},
	}
	dir := t.TempDir()
	n, err := Copy(fsys, dir)
	if err != nil || n != 2 {
		t.Fatalf("Copy = %d, %v; want 2, nil", n, err)
	}
	for p, f := range fsys {
		got, err := os.ReadFile(filepath.Join(dir, p))
		if err != nil || string(got) != string(f.Data) {
			t.Errorf("%s = %q (err %v), want %q", p, got, err, f.Data)
		}
	}
	if n, err := Copy(fsys, dir); err != nil || n != 2 {
		t.Errorf("second Copy = %d, %v; want 2, nil", n, err)
	}
	blocked := filepath.Join(dir, "README.md", "sub") // parent is a regular file
	if _, err := Copy(fsys, blocked); err == nil {
		t.Error("Copy into a path under a regular file expected error, got nil")
	}
}
