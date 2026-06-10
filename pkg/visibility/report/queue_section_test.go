//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what queueSection이 kind·status 표를 내고 빈 윈도는 고정 한 줄인지 검증
package report

import (
	"strings"
	"testing"
)

func TestQueueSection(t *testing.T) {
	if got := queueSection(nil); got != "no queue intake in this window\n" {
		t.Errorf("empty window line wrong: %q", got)
	}
	got := queueSection([]QueueCount{{Kind: "refresh", Status: "open", Count: 2}})
	if !strings.Contains(got, "| refresh | open | 2 |") {
		t.Errorf("table row missing: %q", got)
	}
}
