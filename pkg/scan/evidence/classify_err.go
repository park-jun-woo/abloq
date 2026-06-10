//ff:func feature=scan type=rule control=sequence topic=evidence
//ff:what 네트워크 오류 → 점검 판정 — DNS 미존재(도메인 사망)는 hard, 타임아웃·접속 거부 등 그 외 오류는 soft
package evidence

import (
	"errors"
	"net"
)

// classifyErr maps a probe's transport error to the check status. A domain
// that no longer resolves is a hard failure (the citation cannot recover by
// retrying); everything else — timeouts, refused connections, TLS trouble —
// is soft and must persist to count toward rot.
func classifyErr(err error) string {
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
		return "hard"
	}
	return "soft"
}
