//ff:func feature=gate type=parser control=iteration dimension=1 topic=lossless
//ff:what ContentLines가 이미지/저작자 표기/섹션 헤딩을 제외한 정규화 본문 라인만 반환하는지 검증
package gate

import "testing"

func TestContentLines(t *testing.T) {
	b := loadGateBlog(t)
	d := ParseArticle(b, "en", "---\ntitle: x\n---\n\n![i](/i)\n*Image: AI generated*\n\nBody One.\n\n## Sources\n\n- src\n\n### Related\n\nTail.\n")
	got := ContentLines(d)
	want := []string{"body one.", "- src", "tail."}
	if len(got) != len(want) {
		t.Fatalf("ContentLines = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("ContentLines[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
