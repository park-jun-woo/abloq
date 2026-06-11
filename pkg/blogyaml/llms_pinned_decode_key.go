//ff:func feature=blogyaml type=parser control=selection
//ff:what pinned 엔트리의 키 1개를 디코드 — title/url/desc/group만 허용, 그 외는 unknown-key(strict-동등) 에러
package blogyaml

import "gopkg.in/yaml.v3"

// decodeLlmsPinnedKey decodes one known pinned-entry key; unknown keys are
// rejected with the same message shape as KnownFields(true).
func decodeLlmsPinnedKey(p *LlmsPinned, key, val *yaml.Node) error {
	switch key.Value {
	case "title":
		return val.Decode(&p.Title)
	case "url":
		return val.Decode(&p.URL)
	case "desc":
		return val.Decode(&p.Desc)
	case "group":
		return val.Decode(&p.Group)
	}
	return llmsUnknownKey(key, "blogyaml.LlmsPinned")
}
