//ff:func feature=gate type=rule control=sequence
//ff:what siblingLangs가 그룹별 존재 언어 목록을 수집 순서대로 만드는지 검증
package gate

import (
	"reflect"
	"testing"
)

func TestSiblingLangs(t *testing.T) {
	arts := []*Article{
		{Lang: "ko", Section: "tech", Slug: "a"},
		{Lang: "en", Section: "tech", Slug: "a"},
		{Lang: "ko", Section: "tech", Slug: "b"},
	}
	sibs := siblingLangs(arts)
	if !reflect.DeepEqual(sibs["tech/a"], []string{"ko", "en"}) {
		t.Errorf("tech/a = %v, want [ko en]", sibs["tech/a"])
	}
	if !reflect.DeepEqual(sibs["tech/b"], []string{"ko"}) {
		t.Errorf("tech/b = %v, want [ko]", sibs["tech/b"])
	}
}
