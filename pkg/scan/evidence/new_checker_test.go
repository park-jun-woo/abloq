//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what NewChecker 기본값(타임아웃·동시성·간격·UA)과 LINKCHECK_HOST_OVERRIDE env 반영 검증
package evidence

import (
	"testing"
	"time"
)

func TestNewChecker(t *testing.T) {
	t.Setenv("LINKCHECK_HOST_OVERRIDE", "http://127.0.0.1:9999")
	c := NewChecker()
	if c.Client.Timeout != 10*time.Second {
		t.Errorf("timeout = %v", c.Client.Timeout)
	}
	if c.Concurrency != 4 || c.DomainDelay != 500*time.Millisecond {
		t.Errorf("limits = (%d, %v)", c.Concurrency, c.DomainDelay)
	}
	if c.UserAgent != "abloqd-linkcheck" {
		t.Errorf("UA = %q", c.UserAgent)
	}
	if c.HostOverride != "http://127.0.0.1:9999" {
		t.Errorf("override = %q", c.HostOverride)
	}
}
