//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what llmsDecodeErrors가 nil은 빈 목록, TypeError는 내부 에러 전개, 그 외 에러는 메시지 1건으로 환원하는지 검증
package blogyaml

import (
	"errors"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLlmsDecodeErrors(t *testing.T) {
	if got := llmsDecodeErrors(nil); got != nil {
		t.Errorf("nil error = %v, want nil", got)
	}
	te := &yaml.TypeError{Errors: []string{"line 1: a", "line 2: b"}}
	if got := llmsDecodeErrors(te); !reflect.DeepEqual(got, te.Errors) {
		t.Errorf("TypeError = %v, want inner errors %v", got, te.Errors)
	}
	plain := errors.New("boom")
	if got := llmsDecodeErrors(plain); !reflect.DeepEqual(got, []string{"boom"}) {
		t.Errorf("plain error = %v, want [boom]", got)
	}
}
