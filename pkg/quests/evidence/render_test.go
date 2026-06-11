//ff:func feature=quest type=generator control=sequence topic=queue
//ff:what Render가 대상 경로·claims·rot_urls 원문과 프로토콜·tasks·context 문서를 포함하는지 검증 (불량 payload는 에러)
package evidence

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
		"evidence quest — en/posts/a",
		"content/en/posts/a.md",
		unsourcedClaim, // the claims JSON rides verbatim
		rotURL,
		"abloq quest evidence submit --key en/posts/a",
		"순서 박제",  // _queue-protocol.md
		"검출 내역",  // tasks.md
		"보강의 본질", // context.md
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
