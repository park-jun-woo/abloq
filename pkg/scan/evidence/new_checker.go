//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what 기본 점검기 생성 — 10초 타임아웃, 호스트 4 동시, 도메인 간격 500ms, UA abloqd-linkcheck, env LINKCHECK_HOST_OVERRIDE
package evidence

import (
	"net/http"
	"os"
	"time"
)

// NewChecker builds the default link checker. The check is a light HEAD (GET
// fallback) at a low cadence, so politeness is the declared User-Agent plus
// the per-domain interval; LINKCHECK_HOST_OVERRIDE points every probe at a
// local stub in tests (no real network).
func NewChecker() *Checker {
	return &Checker{
		Client:       &http.Client{Timeout: 10 * time.Second},
		Concurrency:  4,
		DomainDelay:  500 * time.Millisecond,
		UserAgent:    "abloqd-linkcheck",
		HostOverride: os.Getenv("LINKCHECK_HOST_OVERRIDE"),
	}
}
