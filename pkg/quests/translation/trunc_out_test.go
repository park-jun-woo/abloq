//ff:func feature=quest type=rule control=sequence
//ff:what truncOut 검증 — 출력 공백 정리, 300자 절단, 빈 출력은 에러 문자열 대체
package translation

import (
	"fmt"
	"strings"
	"testing"
)

func TestTruncOut(t *testing.T) {
	if got := truncOut([]byte("  boom \n"), fmt.Errorf("exit 1")); got != "boom" {
		t.Errorf("trimmed = %q", got)
	}
	if got := truncOut(nil, fmt.Errorf("exit 1")); got != "exit 1" {
		t.Errorf("empty out = %q", got)
	}
	long := strings.Repeat("x", 400)
	if got := truncOut([]byte(long), nil); len(got) != 303 || !strings.HasSuffix(got, "...") {
		t.Errorf("long out len = %d", len(got))
	}
}
