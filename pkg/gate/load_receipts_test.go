//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what loadReceipts 케이스 — 파일 없음·깨진 JSON·null은 빈 맵, 저장본 라운드트립 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadReceipts(t *testing.T) {
	dir := t.TempDir()
	if got := loadReceipts(dir); len(got) != 0 {
		t.Errorf("missing file: want empty map, got %v", got)
	}
	writeRepoFile(t, dir, filepath.Join(".abloq", "citation-receipts.json"), "{broken")
	if got := loadReceipts(dir); len(got) != 0 {
		t.Errorf("corrupt file: want empty map, got %v", got)
	}
	writeRepoFile(t, dir, filepath.Join(".abloq", "citation-receipts.json"), "null")
	if loadReceipts(dir) == nil {
		t.Error("null file: want non-nil empty map")
	}
	want := map[string]receipt{"https://x.test/a": {CheckedAt: time.Now().UTC().Truncate(time.Second), Verdict: "ok"}}
	if err := saveReceipts(dir, want); err != nil {
		t.Fatalf("saveReceipts: %v", err)
	}
	got := loadReceipts(dir)
	if len(got) != 1 || got["https://x.test/a"].Verdict != "ok" {
		t.Errorf("roundtrip: got %v, want %v", got, want)
	}
	_ = os.RemoveAll(dir)
}
