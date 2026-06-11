//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — 기준선/제출본 마크다운 쌍으로 Consumption(기준선 부착 Target) 조립, git 불요
package common

import (
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func consFixture(t *testing.T, baseMD, docMD string) *Consumption {
	t.Helper()
	root, _ := writeFixture(t, "content/en/posts/a.md", docMD)
	tgt, _, err := AssembleTarget(root, "content/en/posts/a.md", "en", "posts", "a")
	if err != nil {
		t.Fatalf("AssembleTarget: %v", err)
	}
	tgt.Articles[0].Base = agate.ParseArticle(tgt.Blog, "en", baseMD)
	return &Consumption{Target: tgt, Allowed: map[string]bool{}, QueuedClaims: map[string]bool{}}
}
