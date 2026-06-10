//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what IngestKeys가 gzip이 아닌 로그 객체를 에러로 거부하는지 검증
package cflog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIngestKeysBadGzip(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "E.2026-06-01-12.x.gz"), []byte("not gzip"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := IngestKeys(DirSource{Root: dir}, nil, []string{"E.2026-06-01-12.x.gz"}); err == nil {
		t.Error("non-gzip object accepted")
	}
}
