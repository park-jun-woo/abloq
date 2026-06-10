//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what DirSource.Get이 키의 파일 스트림을 열고 없는 키는 에러인지 검증
package cflog

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestDirGet(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "key.gz"), []byte("body"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	src := DirSource{Root: dir}
	rc, err := src.Get("key.gz")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer rc.Close()
	body, _ := io.ReadAll(rc)
	if string(body) != "body" {
		t.Errorf("body = %q", body)
	}
	if _, err := src.Get("missing.gz"); err == nil {
		t.Errorf("missing key accepted")
	}
}
