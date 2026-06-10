//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what Tokens가 공백·구두점·기호를 경계로 소문자 토큰열을 만드는지 검증
package gate

import (
	"reflect"
	"testing"
)

func TestTokens(t *testing.T) {
	got := Tokens("Hello, World!  (FOO-bar) 한국어.\n100%")
	want := []string{"hello", "world", "foo", "bar", "한국어", "100"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Tokens = %v, want %v", got, want)
	}
	if n := len(Tokens("  ,.;! ")); n != 0 {
		t.Errorf("punct-only: want 0 tokens, got %d", n)
	}
}
