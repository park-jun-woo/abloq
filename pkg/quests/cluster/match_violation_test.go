//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what matchViolation이 큐 지정 룰과 겹치는 첫 위반을 반환하고, 무관 위반만이면 해소, 불량 JSON은 미해소인지 검증
package cluster

import (
	"strings"
	"testing"
)

func TestMatchViolation(t *testing.T) {
	rules := map[string]bool{"no-isolated-post": true}
	raw := `[{"rule":"min-internal-links","detail":"d"},{"rule":"no-isolated-post","detail":"d"}]`
	if rule, remains := matchViolation(raw, rules); !remains || rule != "no-isolated-post" {
		t.Errorf("matched = %s %v", rule, remains)
	}
	unrelated := `[{"rule":"tag-taxonomy","detail":"d"}]`
	if rule, remains := matchViolation(unrelated, rules); remains {
		t.Errorf("unrelated violation matched: %s", rule)
	}
	if rule, remains := matchViolation("not json", rules); !remains || !strings.Contains(rule, "unreadable") {
		t.Errorf("malformed payload must count as unresolved: %s %v", rule, remains)
	}
}
