//ff:func feature=blogyaml type=parser control=iteration dimension=1
//ff:what geo.llms_txt 객체 폼의 키-값 쌍을 순회 디코드 — 키별 에러(미지 키 포함)를 전부 모아 TypeError로 반환
package blogyaml

import "gopkg.in/yaml.v3"

// decodeLlmsTxtMap decodes the object form of geo.llms_txt, collecting every
// per-key error so multiple problems surface in one run (KnownFields parity).
func decodeLlmsTxtMap(s *LlmsTxtSpec, n *yaml.Node) error {
	var errs []string
	for i := 0; i+1 < len(n.Content); i += 2 {
		errs = append(errs, llmsDecodeErrors(decodeLlmsTxtKey(s, n.Content[i], n.Content[i+1]))...)
	}
	if len(errs) > 0 {
		return &yaml.TypeError{Errors: errs}
	}
	return nil
}
