//ff:type feature=scan type=client topic=evidence
//ff:what link rot 점검기 — HTTP 클라이언트, 호스트 동시성 상한, 도메인별 요청 간격, UA, 스텁 호스트 오버라이드
//ff:why 외부 호출은 pkg 내부에 가둔다(func는 net/http 금지) — 테스트·Hurl은 LINKCHECK_HOST_OVERRIDE로 전 URL을 로컬 스텁에 수렴시킨다
package evidence

import (
	"net/http"
	"time"
)

// Checker probes citation URLs. Hosts run in parallel up to Concurrency;
// URLs of one host run sequentially with DomainDelay between probes (the
// per-domain rate limit). HostOverride, when set, rewrites every probe's
// scheme+host to it (local stub routing for tests) while results stay keyed
// on the original URL.
type Checker struct {
	Client       *http.Client
	Concurrency  int
	DomainDelay  time.Duration
	UserAgent    string
	HostOverride string
}
