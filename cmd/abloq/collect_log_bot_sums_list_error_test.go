//ff:func feature=cli type=command control=sequence topic=report
//ff:what collectLogBotSums가 stat은 되지만 나열 불가(읽기 권한 없음)한 로그 디렉토리를 에러로 내는지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestCollectLogBotSumsListError(t *testing.T) {
	b, _, err := blogyaml.Load("../../backend/fixtures/blog/blog.yaml")
	if err != nil {
		t.Fatal(err)
	}
	// A directory that stats fine but cannot be listed (no read permission).
	locked := filepath.Join(t.TempDir(), "locked")
	if err := os.Mkdir(locked, 0o311); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(locked, 0o755) })
	if _, _, err := collectLogBotSums("../../backend/fixtures/blog", locked, "2026-06", b); err == nil {
		t.Skip("running with CAP_DAC_OVERRIDE — list error not reproducible")
	}
}
