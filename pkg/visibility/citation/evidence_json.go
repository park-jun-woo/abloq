//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what 근거 JSON 직렬화 — 매칭 URL 목록 또는 엔진 에러를 {"matched":[...]}/{"error":"..."} 문자열로 (스키마 불변, evidence TEXT 재사용)
package citation

import "encoding/json"

// evidenceJSON serializes one sample's evidence: the matched own-domain URLs
// on success, the engine error otherwise. Marshal cannot fail on these
// shapes, so the fallback branch is unreachable in practice.
func evidenceJSON(matched []string, errMsg string) string {
	if errMsg != "" {
		b, _ := json.Marshal(map[string]string{"error": errMsg})
		return string(b)
	}
	if matched == nil {
		matched = []string{}
	}
	b, _ := json.Marshal(map[string][]string{"matched": matched})
	return string(b)
}
