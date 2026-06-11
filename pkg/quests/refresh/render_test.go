//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render가 대상 경로·발급 근거·제출 명령과 프로토콜·tasks·context 문서를 포함하는지 검증 (불량 payload는 에러)
package refresh

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRender(t *testing.T) {
	root := writeInstance(t)
	items, err := Definition{}.Seed([]string{root})
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	out, err := Definition{}.Render(nil, items[0])
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	for _, want := range []string{
		"refresh quest — en/posts/a",
		"content/en/posts/a.md",
		"lastmod 2026-06-02 exceeded the 90-day freshness window",
		"abloq quest refresh submit --key en/posts/a",
		"순서 박제",       // _queue-protocol.md
		"낡음 진단",       // tasks.md
		"갱신의 본질",      // context.md
	} {
		if !strings.Contains(out, want) {
			t.Errorf("prompt missing %q", want)
		}
	}
	bad := &quest.Item{Key: "x", Payload: []byte("not json")}
	if _, err := (Definition{}).Render(nil, bad); err == nil {
		t.Error("corrupt payload: want error")
	}
}
