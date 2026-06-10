//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what evidenceJSON이 매칭 URL은 {"matched":[...]}, 엔진 에러는 {"error":"..."}, 빈 매칭은 빈 배열로 직렬화하는지 검증
package citation

import "testing"

func TestEvidenceJSON(t *testing.T) {
	got := evidenceJSON([]string{"https://blog.test/a/"}, "")
	if got != `{"matched":["https://blog.test/a/"]}` {
		t.Errorf("matched evidence = %s", got)
	}
	if got := evidenceJSON(nil, ""); got != `{"matched":[]}` {
		t.Errorf("empty evidence = %s", got)
	}
	if got := evidenceJSON([]string{"x"}, "engine exploded"); got != `{"error":"engine exploded"}` {
		t.Errorf("error evidence = %s (error must win)", got)
	}
}
