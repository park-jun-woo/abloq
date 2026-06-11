//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what readQueueFile이 직렬화 파일을 Item으로 읽고 부재·변조 파일은 에러인지 검증
package common

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestReadQueueFile(t *testing.T) {
	dir := t.TempDir()
	it := queueio.Item{Kind: "refresh", Slug: "a", Lang: "en", Section: "posts", Priority: 3,
		Payload: map[string]string{}}
	path := filepath.Join(dir, "q.yaml")
	if err := os.WriteFile(path, queueio.Serialize(it), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := readQueueFile(path)
	if err != nil || got.Slug != "a" || got.Priority != 3 {
		t.Errorf("got = %+v (%v)", got, err)
	}
	if _, err := readQueueFile(filepath.Join(dir, "missing.yaml")); err == nil {
		t.Error("missing file: want error")
	}
	if err := os.WriteFile(path, []byte("tampered\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := readQueueFile(path); err == nil {
		t.Error("tampered file: want error")
	}
}
