//ff:func feature=archive type=client control=sequence
//ff:what existingSet이 (kind, target) 페어를 키로 집합화하고 kind/target 충돌을 구분하는지 검증
package archive

import "testing"

func TestExistingSet(t *testing.T) {
	seen := existingSet([]Existing{
		{Kind: KindWayback, Target: "https://a/"},
		{Kind: KindIndexNow, Target: "https://b/"},
	})
	if !seen[KindWayback+"\n"+"https://a/"] || !seen[KindIndexNow+"\n"+"https://b/"] {
		t.Errorf("seen = %v, want both recorded pairs", seen)
	}
	if seen[KindWayback+"\n"+"https://b/"] {
		t.Error("kind/target must not cross-match")
	}
	if len(existingSet(nil)) != 0 {
		t.Error("empty input must produce an empty set")
	}
}
