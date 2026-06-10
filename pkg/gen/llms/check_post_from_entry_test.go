//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what 디렉토리에서 이름이 일치하는 엔트리를 찾아 postFromEntry 결과(ok/slug/title)를 검증
package llms

import (
	"os"
	"testing"
)

func checkPostFromEntry(t *testing.T, dir, entryName string, wantOK bool, wantSlug, wantTitle string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	for _, entry := range entries {
		if entry.Name() != entryName {
			continue
		}
		p, ok := postFromEntry(dir, "ko", "opinion", entry)
		checkPostResult(t, p, ok, wantOK, wantSlug, wantTitle)
		return
	}
	t.Fatalf("entry %q not found in %s", entryName, dir)
}
