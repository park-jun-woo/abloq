//ff:func feature=gen type=generator control=sequence
//ff:what Build 산출물 1개를 testdata/golden/want/의 스냅샷 파일과 바이트 비교 검증
package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func checkGoldenOutput(t *testing.T, dir string, o Output) {
	t.Helper()
	wantPath := filepath.Join(dir, "want", filepath.Base(o.Path))
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("read snapshot %s: %v", wantPath, err)
	}
	if string(o.Data) != string(want) {
		line, w, g := firstDiffLine(want, o.Data)
		t.Errorf("%s drifts from snapshot at line %d: want %q, got %q", o.Path, line, w, g)
	}
}
