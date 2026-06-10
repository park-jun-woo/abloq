//ff:func feature=gate type=output control=sequence topic=diagnostics
//ff:what trunc가 80바이트 이하를 보존하고 초과분을 말줄임표로 자르는지 검증
package gate

import (
	"strings"
	"testing"
)

func TestTrunc(t *testing.T) {
	if got := trunc("short"); got != "short" {
		t.Errorf("trunc(short) = %q", got)
	}
	long := strings.Repeat("x", 100)
	got := trunc(long)
	if len(got) != 80+len("…") || !strings.HasSuffix(got, "…") {
		t.Errorf("trunc(long) = %q (len %d)", got, len(got))
	}
}
