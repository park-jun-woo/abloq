//ff:func feature=gate type=rule control=sequence topic=baseline
//ff:what honestLastmodDiags가 큐 도입 저장소에서 미등재 글의 lastmod 갱신을 차단하고 등재 글을 허용하는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHonestLastmodDiags(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromMD(t, b, "en", "tech", "base", "baseline/lastmod-real.md")
	a.Base = artFromMD(t, b, "en", "tech", "base", "baseline/base.md").Doc

	dir := t.TempDir()
	queueDir := filepath.Join(dir, "quests", "queue")
	if err := os.MkdirAll(queueDir, 0o755); err != nil {
		t.Fatal(err)
	}
	tgt := NewTarget(dir, b, []*Article{a})
	diags := honestLastmodDiags(tgt, a)
	checkDiags(t, diags, 1, "honest-lastmod", "not in the freshness queue")

	queueFile := filepath.Join(queueDir, "freshness.yaml")
	if err := os.WriteFile(queueFile, []byte("key: \"en/tech/base\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := honestLastmodDiags(tgt, a); len(got) != 0 {
		t.Errorf("queued article: want 0 diagnostics, got %v", got)
	}
}
