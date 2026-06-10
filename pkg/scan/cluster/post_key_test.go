//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what PostKey가 <section>/<slug>를 그대로 조립하고 섹션이 다르면 키도 다른지 검증
package cluster

import "testing"

func TestPostKey(t *testing.T) {
	if got := PostKey("tech", "hub"); got != "tech/hub" {
		t.Errorf("PostKey = %q, want tech/hub", got)
	}
	if PostKey("tech", "hub") == PostKey("opinion", "hub") {
		t.Error("same slug in different sections must produce different keys")
	}
}
