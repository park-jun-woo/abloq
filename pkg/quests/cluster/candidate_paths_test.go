//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what candidatePaths가 후보마다 플랫/번들 2형 경로를 펼치고 부재는 nil, 불량 JSON은 에러인지 검증
package cluster

import (
	"reflect"
	"testing"
)

func TestCandidatePaths(t *testing.T) {
	got, err := candidatePaths(map[string]string{
		"candidates": `[{"section":"posts","slug":"hub","shared_tags":1,"directions":["in"]}]`}, "en")
	want := []string{"content/en/posts/hub.md", "content/en/posts/hub/index.md"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v (%v)", got, err)
	}
	got, err = candidatePaths(map[string]string{}, "en")
	if err != nil || got != nil {
		t.Errorf("absent candidates must yield nil: %v (%v)", got, err)
	}
	if _, err := candidatePaths(map[string]string{"candidates": "not json"}, "en"); err == nil {
		t.Error("malformed candidates: want error")
	}
}
