//ff:func feature=queueio type=parser control=sequence
//ff:what liftKeys가 payload의 keys JSON을 1급 목록으로 분리·제거하고 부재는 nil, 불량 JSON은 keys만 버리는지 검증
package queueio

import (
	"reflect"
	"testing"
)

func TestLiftKeys(t *testing.T) {
	keys, rest := liftKeys(map[string]string{"keys": `["ko/t/a","en/t/a"]`, "lastmod": "2026-06-01"})
	if !reflect.DeepEqual(keys, []string{"ko/t/a", "en/t/a"}) {
		t.Errorf("keys = %v", keys)
	}
	if _, dup := rest["keys"]; dup || rest["lastmod"] != "2026-06-01" {
		t.Errorf("rest = %v", rest)
	}
	orig := map[string]string{"lastmod": "2026-06-01"}
	keys, rest = liftKeys(orig)
	if keys != nil || !reflect.DeepEqual(rest, orig) {
		t.Errorf("absent keys: %v / %v", keys, rest)
	}
	keys, rest = liftKeys(map[string]string{"keys": "not json"})
	if keys != nil {
		t.Errorf("malformed keys must yield nil: %v", keys)
	}
	if _, dup := rest["keys"]; dup {
		t.Error("malformed keys must still be removed from the copy")
	}
}
