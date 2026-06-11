//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render가 대상 경로·violations·candidates 원문과 프로토콜·tasks·context 문서를 포함하는지 검증 (불량 payload는 에러)
package cluster

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
		"cluster quest — en/posts/thin",
		"content/en/posts/thin.md",
		"no-isolated-post",
		`"slug":"hub"`,
		"abloq quest cluster submit --key en/posts/thin",
		"순서 박제",   // _queue-protocol.md
		"위반 확인",   // tasks.md
		"큐레이션의 본질", // context.md
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
