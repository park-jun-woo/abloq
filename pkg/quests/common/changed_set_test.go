//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what ChangedSet이 수정·untracked를 모두 잡고 클린 트리는 빈 집합, 비git 디렉토리는 에러인지 검증
package common

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestChangedSet(t *testing.T) {
	root, abs := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	gitFixture(t, root)
	got, err := ChangedSet(root)
	if err != nil {
		t.Fatalf("ChangedSet: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("clean tree: want empty set, got %v", got)
	}
	if err := os.WriteFile(abs, []byte(fixtureArticleMD+"\nMore.\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "stray.md"), []byte("untracked\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err = ChangedSet(root)
	if err != nil {
		t.Fatalf("ChangedSet dirty: %v", err)
	}
	want := []string{"content/en/posts/a.md", "stray.md"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("changed = %v, want %v (untracked included)", got, want)
	}
	if _, err := ChangedSet(t.TempDir()); err == nil {
		t.Error("non-git dir: want error")
	}
}
