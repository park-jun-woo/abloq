//ff:func feature=blogyaml type=parser control=iteration dimension=1 topic=diagnostics
//ff:what yamlErrorDiags가 단일 에러는 1건, yaml.TypeError는 메시지당 1건의 진단으로 변환하는지 검증
package blogyaml

import (
	"errors"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYamlErrorDiags(t *testing.T) {
	cases := []struct {
		name      string
		err       error
		wantDiags int
	}{
		{"plain error", errors.New("yaml: line 2: did not find expected key"), 1},
		{"type error multi", &yaml.TypeError{Errors: []string{
			"line 3: field a not found in type blogyaml.Blog",
			"line 5: field b not found in type blogyaml.Blog",
		}}, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { checkYamlErrorDiags(t, tc.err, tc.wantDiags) })
	}
}
