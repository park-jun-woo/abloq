//ff:func feature=quest type=parser control=sequence
//ff:what Prepare 에러 경로 검증 — 제출 JSON 불량, 대상 글 불일치, hugo 부재, 번역 파일 부재, 원문 lastmod 부재 (전부 try 미소진 에러)
package translation

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareErrors(t *testing.T) {
	orig := hugoLook
	hugoLook = func(string) (string, error) { return "/usr/bin/true", nil }
	defer func() { hugoLook = orig }()
	root := writeInstance(t)
	origin, ko := passPair()
	originPath := writeFile(t, root, "content/en/posts/fixture.md", origin)
	items, err := Definition{}.Seed([]string{originPath})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	it := items[0]
	raw := []byte(`{"article":"content/ko/posts/fixture.md"}`)
	if _, _, err := (Definition{}).Prepare(nil, it, []byte("not json")); err == nil {
		t.Error("bad JSON: want error")
	}
	if _, _, err := (Definition{}).Prepare(nil, it, []byte(`{"article":"content/ko/posts/other.md"}`)); err == nil {
		t.Error("article mismatch: want error")
	}
	if _, _, err := (Definition{}).Prepare(nil, it, raw); err == nil {
		t.Error("translation file absent: want error")
	}
	writeFile(t, root, "content/ko/posts/fixture.md", ko)
	hugoLook = func(string) (string, error) { return "", fmt.Errorf("not found") }
	if _, _, err := (Definition{}).Prepare(nil, it, raw); err == nil || !strings.Contains(err.Error(), "hugo") {
		t.Errorf("hugo absent: err = %v, want hugo lookup error", err)
	}
	hugoLook = func(string) (string, error) { return "/usr/bin/true", nil }
	writeFile(t, root, "content/en/posts/fixture.md", removeLine(origin, "lastmod:"))
	if _, _, err := (Definition{}).Prepare(nil, it, raw); err == nil || !strings.Contains(err.Error(), "lastmod") {
		t.Errorf("origin lastmod missing: err = %v, want front-matter-schema pairing error", err)
	}
	writeFile(t, root, "content/en/posts/fixture.md", removeLine(origin, "date:"))
	if _, _, err := (Definition{}).Prepare(nil, it, raw); err == nil || !strings.Contains(err.Error(), "date") {
		t.Errorf("origin date missing: err = %v, want front-matter-schema pairing error", err)
	}
	if err := os.Remove(originPath); err != nil {
		t.Fatal(err)
	}
	if _, _, err := (Definition{}).Prepare(nil, it, raw); err == nil || !strings.Contains(err.Error(), "origin") {
		t.Errorf("origin unreadable: err = %v, want load error", err)
	}
	if _, _, err := (Definition{}).Prepare(nil, &quest.Item{Payload: []byte("broken")}, raw); err == nil {
		t.Error("broken payload: want decode error")
	}
}
