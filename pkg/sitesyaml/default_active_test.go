//ff:func feature=sitesyaml type=parser control=sequence
//ff:what defaultActive가 active 키 부재 항목만 true로 채우고 명시 false는 보존하는지 검증
package sitesyaml

import "testing"

func TestDefaultActive(t *testing.T) {
	s := Sites{Sites: []Site{
		{Name: "a"},                // no active key in the file
		{Name: "b", Active: false}, // explicit active: false
		{Name: "c", Active: true},  // explicit active: true
	}}
	idx := lineIndex{
		"sites[1].active": 6,
		"sites[2].active": 9,
	}
	defaultActive(&s, idx)
	if !s.Sites[0].Active {
		t.Error("absent key must default to true")
	}
	if s.Sites[1].Active {
		t.Error("explicit false must stay false")
	}
	if !s.Sites[2].Active {
		t.Error("explicit true must stay true")
	}
}
