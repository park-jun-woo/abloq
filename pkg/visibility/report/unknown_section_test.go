//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what unknownSection이 UA·히트 표를 내고 빈 목록은 none인지 검증
package report

import (
	"strings"
	"testing"
)

func TestUnknownSection(t *testing.T) {
	if got := unknownSection(nil); got != "none\n" {
		t.Errorf("empty list must read none: %q", got)
	}
	got := unknownSection([]UnknownBot{{UA: "PetalBot", Hits: 3}})
	if !strings.Contains(got, "| PetalBot | 3 |") {
		t.Errorf("table row missing: %q", got)
	}
}
