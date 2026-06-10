//ff:func feature=gate type=rule control=sequence
//ff:what queueAllows가 큐 미도입=통과, 등재=통과, 미등재=거부를 판정하는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueueAllows(t *testing.T) {
	dir := t.TempDir()
	if !queueAllows(dir, "en/tech/a") {
		t.Error("no queue dir: want allowed (skip)")
	}
	queueDir := filepath.Join(dir, "quests", "queue")
	if err := os.MkdirAll(queueDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if queueAllows(dir, "en/tech/a") {
		t.Error("empty queue: want denied")
	}
	if err := os.MkdirAll(filepath.Join(queueDir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(dir, "nope"), filepath.Join(queueDir, "dangling.yaml")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(queueDir, "q.yaml"), []byte("items:\n  - en/tech/a\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !queueAllows(dir, "en/tech/a") {
		t.Error("queued key: want allowed")
	}
	if queueAllows(dir, "en/tech/zzz") {
		t.Error("unqueued key: want denied")
	}
}
