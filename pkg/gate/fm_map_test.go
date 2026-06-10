//ff:func feature=gate type=parser control=sequence
//ff:what fmMap이 front matter를 맵으로 디코드하고 깨진 YAML에 false를 반환하는지 검증
package gate

import "testing"

func TestFMMap(t *testing.T) {
	m, ok := fmMap("title: \"X\"\ntags: [a, b]\n")
	if !ok {
		t.Fatal("want ok for valid front matter")
	}
	if m["title"] != "X" {
		t.Errorf("title = %v, want X", m["title"])
	}
	if tags, _ := m["tags"].([]any); len(tags) != 2 {
		t.Errorf("tags = %v, want 2 entries", m["tags"])
	}
	if _, ok := fmMap("title: [unclosed"); ok {
		t.Error("want !ok for broken YAML")
	}
}
