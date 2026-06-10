//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what IntervalFromEnv가 CITATION_INTERVAL_MS를 읽고 기본 1초, 0 허용, 음수·비숫자는 기본값인지 검증
package citation

import (
	"testing"
	"time"
)

func TestIntervalFromEnv(t *testing.T) {
	t.Setenv("CITATION_INTERVAL_MS", "")
	if got := IntervalFromEnv(); got != time.Second {
		t.Errorf("default = %v, want 1s", got)
	}
	t.Setenv("CITATION_INTERVAL_MS", "0")
	if got := IntervalFromEnv(); got != 0 {
		t.Errorf("zero = %v, want 0", got)
	}
	t.Setenv("CITATION_INTERVAL_MS", "250")
	if got := IntervalFromEnv(); got != 250*time.Millisecond {
		t.Errorf("250ms = %v", got)
	}
	t.Setenv("CITATION_INTERVAL_MS", "-5")
	if got := IntervalFromEnv(); got != time.Second {
		t.Errorf("negative = %v, want default", got)
	}
	t.Setenv("CITATION_INTERVAL_MS", "abc")
	if got := IntervalFromEnv(); got != time.Second {
		t.Errorf("non-numeric = %v, want default", got)
	}
}
