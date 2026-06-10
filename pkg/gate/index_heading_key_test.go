//ff:func feature=gate type=frame control=sequence
//ff:what indexHeadingKey가 언어별 맵을 생성하며 정규화 키로 등록하는지 검증
package gate

import "testing"

func TestIndexHeadingKey(t *testing.T) {
	byLang := map[string]map[string]string{}
	indexHeadingKey(byLang, "sources", map[string]string{"en": "  Sources ", "ko": "출처"})
	indexHeadingKey(byLang, "related", map[string]string{"en": "Related"})
	if byLang["en"]["sources"] != "sources" {
		t.Errorf("en sources = %q, want normalized registration", byLang["en"]["sources"])
	}
	if byLang["en"]["related"] != "related" {
		t.Errorf("en related missing: %v", byLang["en"])
	}
	if byLang["ko"]["출처"] != "sources" {
		t.Errorf("ko = %v", byLang["ko"])
	}
}
