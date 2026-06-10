//ff:type feature=gate type=schema topic=evidence
//ff:what 인용 검증 영수증 1건 — 검증 시각과 판정(ok/broken/meta-mismatch), 24h TTL 내 재검증 생략의 근거
package gate

import "time"

// citationTTL is how long a verification receipt stays valid: the same URL
// is not re-verified within this window.
const citationTTL = 24 * time.Hour

// receipt is one cached citation-verification result. Verdict is "ok",
// "broken" or "meta-mismatch" — transient "retry" results are never cached.
type receipt struct {
	CheckedAt time.Time `json:"checked_at"`
	Verdict   string    `json:"verdict"`
	Detail    string    `json:"detail,omitempty"`
}
