//ff:func feature=quest type=frame control=sequence topic=queue
//ff:what Cons가 자기 자신을 반환해 임베드 승격으로 ConsCarrier 계약을 충족하는지 검증
package common

import "testing"

func TestCons(t *testing.T) {
	c := &Consumption{}
	if c.Cons() != c {
		t.Error("Cons must return the receiver")
	}
	var carrier ConsCarrier = struct{ *Consumption }{c}
	if carrier.Cons() != c {
		t.Error("embedding must satisfy ConsCarrier by promotion")
	}
}
