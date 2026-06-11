//ff:func feature=quest type=parser control=sequence
//ff:what Prepare 검증 — 제출 JSON 디코드, 단일 글 Target(Base nil) 조립, match 미출현 산출, REVIEW·로그 본문 적재
package writing

import "testing"

func TestPrepare(t *testing.T) {
	root := writeInstance(t)
	art, ins := passFixtures()
	specPath := writeFile(t, root, "content/en/posts/hello.insight.yaml", ins)
	// drop the bravo anchor so claim c2 lands in match-missing
	writeFile(t, root, "content/en/posts/hello.md", removeLine(art, "bravo"))
	writeFile(t, root, "quests/writing/logs/hello.md", "log entry\n")
	writeFile(t, root, "quests/writing/reviews/hello.md", "reviewer: ctx-2\n- c2: excluded — out of scope\n")
	items, err := Definition{}.Seed([]string{specPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	raw := []byte(`{"article":"content/en/posts/hello.md","worklog":"quests/writing/logs/hello.md","review":"quests/writing/reviews/hello.md"}`)
	ctx, short, err := Definition{}.Prepare(nil, items[0], raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if short != nil {
		t.Fatalf("short verdict = %v, want nil", short)
	}
	sub := ctx.Submission.(*Submission)
	if len(sub.Target.Articles) != 1 || sub.Target.Articles[0].Base != nil {
		t.Errorf("target: %d article(s), Base=%v — want 1 with Base nil", len(sub.Target.Articles), sub.Target.Articles[0].Base)
	}
	if len(sub.Missing) != 1 || sub.Missing[0].ID != "c2" {
		t.Errorf("Missing = %v, want [c2]", sub.Missing)
	}
	if sub.Review == "" || sub.Worklog == "" {
		t.Errorf("review/worklog not loaded: %q %q", sub.Review, sub.Worklog)
	}
}
