//ff:func feature=queueio type=client control=sequence
//ff:what PublishFile이 디렉토리 경로의 일반 파일(mkdir 실패)과 디렉토리 타겟(write 실패)을 각각 에러로 내는지 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPublishFileWriteErrors(t *testing.T) {
	cfg := bareFixture(t)
	if err := ensureClone(cfg); err != nil {
		t.Fatal(err)
	}
	// A path component that is a regular file fails the mkdir.
	if _, err := PublishFile(cfg, "README.md/x.md", []byte("x")); err == nil {
		t.Error("a file in the directory path must error")
	}
	// A target that is a directory fails the write.
	if err := os.MkdirAll(filepath.Join(cfg.Workdir, "reports", "dir.md"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := PublishFile(cfg, "reports/dir.md", []byte("x")); err == nil {
		t.Error("a directory target must error")
	}
}
