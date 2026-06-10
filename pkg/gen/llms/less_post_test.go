//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what lessPost가 언어 랭크→섹션 랭크→날짜 내림차순→slug 오름차순으로 비교하는지 검증
package llms

import "testing"

func TestLessPost(t *testing.T) {
	cases := []struct {
		name string
		a, b Post
		want bool
	}{
		{"lang rank wins", Post{Lang: "ko"}, Post{Lang: "en"}, true},
		{"section rank wins", Post{Lang: "ko", Section: "opinion"}, Post{Lang: "ko", Section: "tech"}, true},
		{"newer date first", Post{Lang: "ko", Date: "2026-03-01"}, Post{Lang: "ko", Date: "2026-01-01"}, true},
		{"older date later", Post{Lang: "ko", Date: "2026-01-01"}, Post{Lang: "ko", Date: "2026-03-01"}, false},
		{"slug tiebreak", Post{Lang: "ko", Slug: "a"}, Post{Lang: "ko", Slug: "b"}, true},
		{"equal posts", Post{Lang: "ko", Slug: "a"}, Post{Lang: "ko", Slug: "a"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkLessPost(t, tc.a, tc.b, tc.want) })
	}
}
