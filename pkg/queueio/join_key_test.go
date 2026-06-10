//ff:func feature=queueio type=generator control=sequence
//ff:what JoinKey가 <lang>/<section>/<slug> 원문 키를 조립하는지 검증
package queueio

import "testing"

func TestJoinKey(t *testing.T) {
	if got := JoinKey("ko", "tech", "post-a"); got != "ko/tech/post-a" {
		t.Errorf("want ko/tech/post-a, got %s", got)
	}
}
