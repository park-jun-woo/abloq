//ff:func feature=insight type=parser control=sequence
//ff:what insight.yaml 파일을 읽어 strict 파싱 + 검증 — (Insight, 에러 진단, 경고 진단, IO 에러)
package insight

import (
	"os"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// Load reads, parses and validates an insight.yaml file. The error return is
// for IO failures only; schema problems come back as diagnostics.
func Load(path string) (*Insight, []blogyaml.Diagnostic, []blogyaml.Diagnostic, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}
	ins, diags := Parse(path, data)
	if len(diags) > 0 {
		return nil, diags, nil, nil
	}
	errs, warns := Validate(path, ins, claimLines(data))
	if len(errs) > 0 {
		return nil, errs, warns, nil
	}
	return ins, nil, warns, nil
}
