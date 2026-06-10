//ff:func feature=gate type=parser control=sequence
//ff:what front matter 원문을 비-strict YAML 맵으로 디코드 — 타입 검증(front-matter-schema)의 입력, 깨지면 false
package gate

import "gopkg.in/yaml.v3"

// fmMap decodes the raw front matter into a generic map for type checks.
func fmMap(fm string) (map[string]any, bool) {
	m := map[string]any{}
	if err := yaml.Unmarshal([]byte(fm), &m); err != nil {
		return nil, false
	}
	return m, true
}
