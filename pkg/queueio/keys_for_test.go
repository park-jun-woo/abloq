//ff:func feature=queueio type=generator control=sequence
//ff:what KeysFor가 선언 순서 그대로 전 언어 조인 키를 만들고 빈 언어 목록은 nil인지 검증
package queueio

import (
	"reflect"
	"testing"
)

func TestKeysFor(t *testing.T) {
	got := KeysFor([]string{"ko", "en"}, "tech", "post-a")
	want := []string{"ko/tech/post-a", "en/tech/post-a"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("KeysFor = %v, want %v", got, want)
	}
	if KeysFor(nil, "tech", "post-a") != nil {
		t.Error("empty language list must yield nil")
	}
}
