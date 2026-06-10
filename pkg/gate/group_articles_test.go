//ff:func feature=gate type=rule control=sequence
//ff:what groupArticles가 섹션/어간 키로 언어 버전을 묶는지 검증
package gate

import "testing"

func TestGroupArticles(t *testing.T) {
	arts := []*Article{
		{Lang: "ko", Section: "tech", Slug: "a"},
		{Lang: "en", Section: "tech", Slug: "a"},
		{Lang: "ko", Section: "tech", Slug: "b"},
	}
	groups := groupArticles(arts)
	if len(groups) != 2 {
		t.Fatalf("want 2 groups, got %v", groups)
	}
	if len(groups["tech/a"]) != 2 || len(groups["tech/b"]) != 1 {
		t.Errorf("group sizes wrong: %v", groups)
	}
}
