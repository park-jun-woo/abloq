//ff:func feature=scan type=rule control=sequence topic=evidence
//ff:what classifyErr 케이스 — DNS 미존재(래핑 포함)는 hard, 일반 네트워크 오류는 soft
package evidence

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

func TestClassifyErr(t *testing.T) {
	dead := &net.DNSError{Err: "no such host", Name: "gone.example", IsNotFound: true}
	if got := classifyErr(fmt.Errorf("Head %q: %w", "https://gone.example/x", dead)); got != "hard" {
		t.Errorf("dead domain = %q, want hard", got)
	}
	transient := &net.DNSError{Err: "server misbehaving", Name: "x.example", IsTimeout: true}
	if got := classifyErr(transient); got != "soft" {
		t.Errorf("dns timeout = %q, want soft", got)
	}
	if got := classifyErr(errors.New("connection refused")); got != "soft" {
		t.Errorf("generic error = %q, want soft", got)
	}
}
