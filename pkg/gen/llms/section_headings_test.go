//ff:func feature=gen type=generator control=sequence
//ff:what sectionHeadings가 정렬된 글에서 출현할 그룹 헤딩 텍스트 집합(라벨·다언어 접두 적용, 중복 제거)을 산출하는지 검증
package llms

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestSectionHeadings(t *testing.T) {
	b := &blogyaml.Blog{}
	b.Geo.LlmsTxt.SectionLabels = map[string]string{"opinion": "Concept"}
	sorted := []Post{
		{Lang: "ko", Section: "opinion"},
		{Lang: "ko", Section: "opinion"},
		{Lang: "en", Section: "tech"},
	}
	got := sectionHeadings(b, sorted, true)
	want := map[string]bool{"ko/Concept": true, "en/tech": true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("multi headings = %v, want %v", got, want)
	}
	got = sectionHeadings(b, sorted, false)
	want = map[string]bool{"Concept": true, "tech": true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("single-scope headings = %v, want %v", got, want)
	}
}
