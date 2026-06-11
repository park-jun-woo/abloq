//ff:func feature=quest type=rule control=sequence topic=queue
//ff:what queue-scope가 허용 집합 밖 변경(다른 글·큐 파일 삭제 포함)을 FAIL로 잡고 범위 내 변경은 통과하는지 검증
package common

import (
	"strings"
	"testing"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
)

type consCarrierStub struct{ c *Consumption }

func (s consCarrierStub) Cons() *Consumption { return s.c }

func TestRuleQueueScope(t *testing.T) {
	r := RuleQueueScope()
	if r.Meta.ID != "queue-scope" || r.Meta.Level != rgate.LevelFail {
		t.Fatalf("Meta = %+v", r.Meta)
	}
	allowed := AllowedPaths([]string{"en/posts/a"})
	c := &Consumption{Allowed: allowed, Changed: []string{"content/en/posts/a.md"}}
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("in-scope change fired")
	}
	c.Changed = []string{"content/en/posts/a.md", "content/en/posts/other.md", "blog.yaml"}
	fired, fact := r.Check(rgate.Context{Submission: consCarrierStub{c}})
	if !fired || !strings.Contains(fact.Actual, "out-of-scope") || !strings.Contains(fact.Actual, "외 1건") {
		t.Errorf("out-of-scope changes: fired=%v fact=%+v", fired, fact)
	}
	// Deleting the queue file before the gate is itself out of scope.
	c.Changed = []string{"quests/queue/refresh-en-posts-a.yaml"}
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); !fired {
		t.Error("queue-file change must fire (deletion is the post-gate signal)")
	}
	c.Changed = nil
	if fired, _ := r.Check(rgate.Context{Submission: consCarrierStub{c}}); fired {
		t.Error("empty change set fired")
	}
}
