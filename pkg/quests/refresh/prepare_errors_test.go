//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Prepare 에러 경로 — 불량 JSON·대상 글 불일치·비git 인스턴스(기준선 불가)가 try 미소진 에러인지 검증
package refresh

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPrepareErrors(t *testing.T) {
	root := writeInstance(t)
	items, err := Definition{}.Seed([]string{root})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	it := items[0]
	if _, _, err := (Definition{}).Prepare(nil, it, []byte("not json")); err == nil {
		t.Error("bad JSON: want error")
	}
	if _, _, err := (Definition{}).Prepare(nil, it, []byte(`{"article":"content/en/posts/other.md"}`)); err == nil {
		t.Error("mismatched article: want error")
	}
	// Destroying the repository makes the baseline impossible — Prepare error,
	// never a silent Base-nil fallback.
	if err := os.RemoveAll(filepath.Join(root, ".git")); err != nil {
		t.Fatal(err)
	}
	if _, _, err := (Definition{}).Prepare(nil, it, []byte(`{"article":"content/en/posts/a.md"}`)); err == nil {
		t.Error("non-git instance: want error (no Base-nil fallback)")
	}
}
