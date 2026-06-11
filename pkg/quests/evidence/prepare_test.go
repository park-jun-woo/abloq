//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what Prepare가 payload claims 해시 집합·rot URL 목록을 적재하고 기준선 Target을 조립하는지 검증 (불일치 글·불량 JSON은 에러)
package evidence

import (
	"testing"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func TestPrepare(t *testing.T) {
	root := writeInstance(t)
	ctx := subWith(t, root)
	sub := ctx.Submission.(*Submission)
	if !sub.QueuedClaims[agate.HashText(unsourcedClaim)] || len(sub.QueuedClaims) != 1 {
		t.Errorf("queued claims = %v", sub.QueuedClaims)
	}
	if len(sub.RotURLs) != 1 || sub.RotURLs[0] != rotURL {
		t.Errorf("rot urls = %v", sub.RotURLs)
	}
	if sub.Target.Articles[0].Base == nil {
		t.Fatal("Base must be attached")
	}
	items, err := Definition{}.Seed([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := (Definition{}).Prepare(nil, items[0], []byte("nope")); err == nil {
		t.Error("bad JSON: want error")
	}
	if _, _, err := (Definition{}).Prepare(nil, items[0], []byte(`{"article":"content/en/posts/b.md"}`)); err == nil {
		t.Error("mismatched article: want error")
	}
}
