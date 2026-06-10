//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what Citations가 코드 펜스 안 링크를 제외하고 front matter 보정된 파일 라인으로 추출하는지 검증
package gate

import "testing"

func TestCitations(t *testing.T) {
	b := loadGateBlog(t)
	content := "---\ntitle: x\n---\n\nSee [a](https://x.test/a).\n\n```\n[b](https://x.test/b)\n```\n\nAnd [c](https://x.test/c).\n"
	got := Citations(ParseArticle(b, "en", content))
	if len(got) != 2 {
		t.Fatalf("want 2 citations (fenced link excluded), got %+v", got)
	}
	if got[0].URL != "https://x.test/a" || got[0].Label != "a" || got[0].Line != 5 {
		t.Errorf("first citation = %+v, want a at file line 5", got[0])
	}
	if got[1].URL != "https://x.test/c" || got[1].Line != 11 {
		t.Errorf("second citation = %+v, want c at file line 11", got[1])
	}
}
