//ff:func feature=quest type=parser control=sequence
//ff:what Prepare 에러 경로 검증 — 제출 JSON 불량, 대상 글 불일치(치즈 방어), 대상 글 파일 부재
package writing

import "testing"

func TestPrepareErrors(t *testing.T) {
	root := writeInstance(t)
	_, ins := passFixtures()
	specPath := writeFile(t, root, "content/en/posts/hello.insight.yaml", ins)
	items, err := Definition{}.Seed([]string{specPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	it := items[0]
	if _, _, err := (Definition{}).Prepare(nil, it, []byte("not json")); err == nil {
		t.Error("bad JSON: want error, got nil")
	}
	other := []byte(`{"article":"content/en/posts/other.md"}`)
	if _, _, err := (Definition{}).Prepare(nil, it, other); err == nil {
		t.Error("article mismatch: want error, got nil")
	}
	missing := []byte(`{"article":"content/en/posts/hello.md"}`)
	if _, _, err := (Definition{}).Prepare(nil, it, missing); err == nil {
		t.Error("article file absent: want error, got nil")
	}
}
