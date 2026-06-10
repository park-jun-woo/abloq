//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what 엔진 호출 간 대기시간을 CITATION_INTERVAL_MS에서 읽기 — 기본 1000ms, 0 허용 (엔진별 rate limit 보호)
package citation

import (
	"os"
	"strconv"
	"time"
)

// IntervalFromEnv reads the per-engine inter-call throttle from
// CITATION_INTERVAL_MS (default 1000ms; tests and stubs set 0). The CLI and
// the backend func share it.
func IntervalFromEnv() time.Duration {
	if v := os.Getenv("CITATION_INTERVAL_MS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			return time.Duration(n) * time.Millisecond
		}
	}
	return time.Second
}
