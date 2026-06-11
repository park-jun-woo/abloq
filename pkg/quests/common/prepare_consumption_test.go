//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what PrepareConsumption이 기준선 Target·변경 집합·허용 집합을 채우고 비git 인스턴스는 에러인지 검증
package common

import (
	"os"
	"testing"
)

func TestPrepareConsumption(t *testing.T) {
	root, abs := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	gitFixture(t, root)
	if err := os.WriteFile(abs, []byte(fixtureArticleMD+"\nMore body.\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	p := QueuePayload{Root: root, Article: "content/en/posts/a.md",
		Lang: "en", Section: "posts", Slug: "a", Keys: []string{"en/posts/a"}}
	c, body, err := PrepareConsumption(p)
	if err != nil {
		t.Fatalf("PrepareConsumption: %v", err)
	}
	if c.Target.Articles[0].Base == nil || len(body) == 0 {
		t.Error("baseline target and body bytes must be filled")
	}
	if len(c.Changed) != 1 || c.Changed[0] != "content/en/posts/a.md" {
		t.Errorf("changed = %v", c.Changed)
	}
	if !c.Allowed["content/en/posts/a.md"] {
		t.Errorf("allowed = %v", c.Allowed)
	}
	bare, _ := writeFixture(t, "content/en/posts/a.md", fixtureArticleMD)
	p.Root = bare
	if _, _, err := PrepareConsumption(p); err == nil {
		t.Error("non-git instance: want error")
	}
}
