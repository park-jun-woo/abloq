//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 테스트 헬퍼 — Seed→Prepare 경로로 제출 컨텍스트를 만든다 (실제 submit과 같은 조립)
package evidence

import (
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

func subWith(t *testing.T, root string) rgate.Context {
	t.Helper()
	items, err := Definition{}.Seed([]string{root})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	raw := []byte(`{"article":"content/en/posts/a.md"}`)
	ctx, short, err := Definition{}.Prepare(nil, items[0], raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if short != nil {
		t.Fatalf("unexpected short verdict: %+v", short)
	}
	return ctx
}
