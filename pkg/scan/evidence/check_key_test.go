//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what checkKey가 키 4요소를 개행으로 결합하고 요소가 다르면 키도 달라지는지 검증
package evidence

import "testing"

func TestCheckKey(t *testing.T) {
	k := checkKey("https://a.example/x", "ko", "tech", "post")
	if k != "https://a.example/x\nko\ntech\npost" {
		t.Errorf("key = %q", k)
	}
	if k == checkKey("https://a.example/x", "ko", "tech", "other") {
		t.Error("different slugs must not collide")
	}
}
