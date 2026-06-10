//ff:type feature=scan type=schema topic=evidence
//ff:what 근거 스캔 1회전 결과 — 큐 후보 항목(claims+확정 rot)과 갱신된 인용 점검 상태 전체
package evidence

import "github.com/park-jun-woo/abloq/pkg/queueio"

// Result is one evidence scan pass. Items are the kind=evidence queue
// candidates (unsourced claims plus rot URLs confirmed by ConsecutiveFailures
// >= rotThreshold); Checks is the full updated citation state the backend
// upserts into citation_checks. A stateless caller (the CLI) passes no
// previous checks, so Items can never contain rot — claims only.
type Result struct {
	Items  []queueio.Item
	Checks []Check
}
