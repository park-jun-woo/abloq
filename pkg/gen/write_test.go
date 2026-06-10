//ff:func feature=gen type=generator control=sequence
//ff:what Write가 상위 디렉토리를 만들며 산출물을 기록하고 재기록(2회차)이 no-op인지 검증
package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	dir := t.TempDir()
	outs := []Output{{Path: "static/robots.txt", Data: []byte("User-agent: *\n")}}
	if err := Write(dir, outs); err != nil {
		t.Fatalf("Write: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dir, "static", "robots.txt"))
	if err != nil || string(got) != "User-agent: *\n" {
		t.Fatalf("read back = %q, err %v", got, err)
	}
	if err := Write(dir, outs); err != nil {
		t.Fatalf("second Write: %v", err)
	}
	again, err := os.ReadFile(filepath.Join(dir, "static", "robots.txt"))
	if err != nil || string(again) != string(got) {
		t.Errorf("second Write changed bytes: %q -> %q (err %v)", got, again, err)
	}
}
