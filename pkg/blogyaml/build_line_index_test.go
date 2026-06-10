//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what buildLineIndex가 키 경로(중첩 키/시퀀스 항목)별 소스 라인을 기록하는지 검증
package blogyaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestBuildLineIndex(t *testing.T) {
	src := "site:\n  baseURL: https://example.com\nlanguages:\n  - ko\n  - en\n"
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(src), &doc); err != nil {
		t.Fatalf("yaml.Unmarshal: %v", err)
	}
	idx := buildLineIndex(&doc)
	want := map[string]int{
		"site":         1,
		"site.baseURL": 2,
		"languages":    3,
		"languages[0]": 4,
		"languages[1]": 5,
	}
	for path, line := range want {
		if got, ok := idx[path]; !ok || got != line {
			t.Errorf("idx[%q]: want %d, got %d (present=%v)", path, line, got, ok)
		}
	}
}
