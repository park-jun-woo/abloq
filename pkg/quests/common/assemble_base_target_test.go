//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what AssembleBaseTarget이 git HEAD 원본을 Base로 부착하고, HEAD에 없는 글·비git 저장소는 에러(Base nil 폴백 금지)인지 검증
package common

import (
	"os"
	"strings"
	"testing"
)

func TestAssembleBaseTarget(t *testing.T) {
	root, abs := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	gitFixture(t, root)
	updated := strings.Replace(fixtureArticleMD, "Body text.", "Body text updated meaningfully.", 1)
	if err := os.WriteFile(abs, []byte(updated), 0o644); err != nil {
		t.Fatal(err)
	}
	tgt, body, err := AssembleBaseTarget(root, "content/en/posts/a.md", "en", "posts", "a")
	if err != nil {
		t.Fatalf("AssembleBaseTarget: %v", err)
	}
	a := tgt.Articles[0]
	if a.Base == nil || a.Base == a.Doc {
		t.Fatal("Base must be the parsed HEAD snapshot, distinct from Doc")
	}
	if !strings.Contains(a.Base.Body, "Body text.") || !strings.Contains(string(body), "updated meaningfully") {
		t.Errorf("Base/Doc mixed up: base=%q", a.Base.Body)
	}
	// An article absent from HEAD is a Prepare-level error, never Base nil.
	if err := os.WriteFile(abs+".new.md", []byte(fixtureArticleMD), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := AssembleBaseTarget(root, "content/en/posts/a.md.new.md", "en", "posts", "a.md.new"); err == nil {
		t.Error("article missing from HEAD: want error (no Base-nil fallback)")
	}
	// A non-git instance cannot be consumed at all.
	bare, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	if _, _, err := AssembleBaseTarget(bare, "content/en/posts/a.md", "en", "posts", "a"); err == nil {
		t.Error("non-git instance: want error")
	}
}
