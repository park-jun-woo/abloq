//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what queuedViolations가 violations JSON을 룰 집합으로 디코드하고 부재·공집합·불량 JSON은 에러인지 검증
package cluster

import "testing"

func TestQueuedViolations(t *testing.T) {
	got, err := queuedViolations(map[string]string{
		"violations": `[{"rule":"min-internal-links","detail":"d"},{"rule":"no-isolated-post","detail":"d"}]`})
	if err != nil || len(got) != 2 || !got["min-internal-links"] || !got["no-isolated-post"] {
		t.Errorf("got = %v (%v)", got, err)
	}
	if _, err := queuedViolations(map[string]string{}); err == nil {
		t.Error("absent violations: want error")
	}
	if _, err := queuedViolations(map[string]string{"violations": "[]"}); err == nil {
		t.Error("empty violations: want error")
	}
	if _, err := queuedViolations(map[string]string{"violations": "not json"}); err == nil {
		t.Error("malformed violations: want error")
	}
}
