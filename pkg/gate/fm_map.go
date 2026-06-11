//ff:func feature=gate type=parser control=sequence
//ff:what front matter 원문을 비-strict YAML 맵으로 디코드 — 타입 검증(front-matter-schema)의 입력, 깨지면 false (공개 API)
//ff:why Phase017 번역 퀘스트의 fm-mirror(⑦)·Seed lastmod 선별이 front matter 값 접근을 공유해야 해서 export — 재구현(복제) 대신 단일 출처 유지
package gate

import "gopkg.in/yaml.v3"

// FMMap decodes the raw front matter into a generic map for type checks.
// Exported for the translation quest's front matter reads (Phase017).
func FMMap(fm string) (map[string]any, bool) {
	m := map[string]any{}
	if err := yaml.Unmarshal([]byte(fm), &m); err != nil {
		return nil, false
	}
	return m, true
}
