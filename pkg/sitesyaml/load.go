//ff:func feature=sitesyaml type=parser control=sequence
//ff:what sites.yaml 파일을 읽어 strict 파싱 + 검증까지 수행, (Sites, 진단, IO 에러)를 반환
package sitesyaml

import "os"

// Load reads, parses and validates a sites.yaml file.
// The error return is for IO failures only; schema problems come back as diagnostics.
func Load(path string) (*Sites, []Diagnostic, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	s, idx, diags := Parse(path, data)
	if len(diags) > 0 {
		return nil, diags, nil
	}
	return s, Validate(path, s, idx), nil
}
