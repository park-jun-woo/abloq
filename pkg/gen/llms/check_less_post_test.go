//ff:func feature=gen type=generator control=sequence
//ff:what lessPost 케이스 하나를 ko/en·opinion/tech 랭크로 실행해 비교 결과를 검증
package llms

import "testing"

func checkLessPost(t *testing.T, a, b Post, want bool) {
	t.Helper()
	langRank := rankOf([]string{"ko", "en"})
	sectionRank := rankOf([]string{"opinion", "tech"})
	if got := lessPost(a, b, langRank, sectionRank); got != want {
		t.Errorf("lessPost(%+v, %+v) = %v, want %v", a, b, got, want)
	}
}
