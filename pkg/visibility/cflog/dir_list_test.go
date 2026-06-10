//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what DirSource.List가 정렬된 파일명 키를 prefix·afterKey 필터로 돌려주고 디렉토리는 제외하는지 검증
package cflog

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDirList(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"E.2026-06-01-12.b.gz", "E.2026-06-01-11.a.gz", "other.txt"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	if err := os.Mkdir(filepath.Join(dir, "E.sub"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	src := DirSource{Root: dir}
	got, err := src.List("", "")
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	want := []string{"E.2026-06-01-11.a.gz", "E.2026-06-01-12.b.gz", "other.txt"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("List = %v, want %v", got, want)
	}
	got, err = src.List("E.", "E.2026-06-01-11.a.gz")
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if !reflect.DeepEqual(got, []string{"E.2026-06-01-12.b.gz"}) {
		t.Errorf("filtered List = %v", got)
	}
}
