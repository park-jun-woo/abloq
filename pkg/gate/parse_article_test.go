//ff:func feature=gate type=parser control=sequence
//ff:what 공개 ParseArticle API가 blog.yaml 헤딩 맵 기준으로 섹션을 분류하고 front matter 없는 입력도 처리하는지 검증
package gate

import "testing"

func TestParseArticle(t *testing.T) {
	b := loadGateBlog(t)
	d := ParseArticle(b, "ko", "---\ntitle: x\n---\n\n본문\n\n## 출처\n\n- 근거\n")
	if len(d.Sections) != 1 || d.Sections[0].Key != "sources" {
		t.Fatalf("want one sources section, got %v", d.Sections)
	}
	if d.BodyStart != 4 {
		t.Errorf("BodyStart = %d, want 4", d.BodyStart)
	}
	noFM := ParseArticle(b, "ko", "본문뿐\n")
	if noFM.HasFM || noFM.BodyStart != 1 {
		t.Errorf("want no front matter at BodyStart 1, got %+v", noFM)
	}
}
