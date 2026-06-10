//ff:func feature=gate type=output control=sequence topic=evidence
//ff:what saveReceipts 케이스 — .abloq 디렉토리 자동 생성 기록, 디렉토리 자리에 파일이 있으면 에러
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveReceipts(t *testing.T) {
	dir := t.TempDir()
	if err := saveReceipts(dir, map[string]receipt{"https://x.test/a": {Verdict: "broken", Detail: "HTTP 404"}}); err != nil {
		t.Fatalf("saveReceipts: %v", err)
	}
	if got := loadReceipts(dir)["https://x.test/a"].Detail; got != "HTTP 404" {
		t.Errorf("want detail HTTP 404, got %q", got)
	}
	blocked := t.TempDir()
	if err := os.WriteFile(filepath.Join(blocked, ".abloq"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := saveReceipts(blocked, nil); err == nil {
		t.Error("want error when .abloq exists as a file, got nil")
	}
}
