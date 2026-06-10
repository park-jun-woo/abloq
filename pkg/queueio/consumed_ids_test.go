//ff:func feature=queueio type=rule control=sequence
//ff:what consumedIDs가 파일이 사라진 exported 행만 골라내는지 검증 (정방향 파일명 계산 + 존재 확인)
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConsumedIDs(t *testing.T) {
	dir := t.TempDir()
	kept := Row{ID: 1, Item: Item{Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech"}}
	gone := Row{ID: 2, Item: Item{Kind: "refresh", Slug: "post-b", Lang: "ko", Section: "tech"}}
	if err := os.WriteFile(filepath.Join(dir, Filename(kept.Item)), Serialize(kept.Item), 0o644); err != nil {
		t.Fatal(err)
	}
	ids := consumedIDs(dir, []Row{kept, gone})
	if len(ids) != 1 || ids[0] != 2 {
		t.Errorf("want [2], got %v", ids)
	}
	if got := consumedIDs(dir, nil); len(got) != 0 {
		t.Errorf("no exported rows must yield empty: %v", got)
	}
}
