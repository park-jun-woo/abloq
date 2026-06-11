//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Prepare가 기준선 부착 Target·porcelain 변경 집합·허용 집합을 채운 Submission을 만드는지 검증
package refresh

import (
	"strings"
	"testing"
)

func TestPrepare(t *testing.T) {
	root := writeInstance(t)
	refreshed := strings.Replace(baseArticleMD, "lastmod: 2026-06-02", "lastmod: 2026-06-09", 1)
	writeFile(t, root, "content/en/posts/a.md", refreshed)
	ctx := subWith(t, root)
	sub := ctx.Submission.(*Submission)
	a := sub.Target.Articles[0]
	if a.Base == nil || a.Base == a.Doc {
		t.Fatal("Base must be the parsed HEAD snapshot")
	}
	if len(sub.Changed) != 1 || sub.Changed[0] != "content/en/posts/a.md" {
		t.Errorf("changed = %v", sub.Changed)
	}
	if !sub.Allowed["content/en/posts/a.md"] || sub.Allowed["quests/queue/refresh-en-posts-a.yaml"] {
		t.Errorf("allowed set wrong: %v", sub.Allowed)
	}
	if !strings.Contains(ctx.Source, "lastmod: 2026-06-09") {
		t.Error("Source must be the working-tree article bytes")
	}
}
