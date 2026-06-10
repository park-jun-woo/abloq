//ff:func feature=blogyaml type=parser control=sequence
//ff:what blog.yaml 파일을 읽어 strict 파싱 + 검증까지 수행, (Blog, 진단, IO 에러)를 반환
package blogyaml

import "os"

// Load reads, parses and validates a blog.yaml file.
// The error return is for IO failures only; schema problems come back as diagnostics.
func Load(path string) (*Blog, []Diagnostic, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	b, idx, diags := Parse(path, data)
	if len(diags) > 0 {
		return nil, diags, nil
	}
	return b, Validate(path, b, idx), nil
}
