//ff:func feature=gate type=frame control=sequence
//ff:what buildHeadingIndex가 order 순위와 언어별 정규화 역색인을 만드는지 검증
package gate

import "testing"

func TestBuildHeadingIndex(t *testing.T) {
	hi := buildHeadingIndex(loadGateBlog(t))
	if hi.rank["image"] != 0 || hi.rank["body"] != 2 || hi.rank["sources"] != 5 {
		t.Errorf("rank = %v, want structure.order indexes", hi.rank)
	}
	if hi.byLang["en"]["further reading"] != "further" {
		t.Errorf("byLang en = %v, want normalized lookup", hi.byLang["en"])
	}
	if hi.byLang["ko"]["변경 이력"] != "changelog" {
		t.Errorf("byLang ko = %v, want changelog", hi.byLang["ko"])
	}
}
