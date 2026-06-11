//ff:func feature=blogyaml type=parser control=sequence topic=diagnostics
//ff:what 키 디코드 에러 1건을 메시지 목록으로 환원 — nil은 빈 목록, TypeError는 내부 에러 전개, 그 외는 메시지 1건
package blogyaml

import (
	"errors"

	"gopkg.in/yaml.v3"
)

// llmsDecodeErrors flattens one per-key decode error into its messages so the
// map loop can aggregate everything into a single TypeError.
func llmsDecodeErrors(err error) []string {
	if err == nil {
		return nil
	}
	var te *yaml.TypeError
	if errors.As(err, &te) {
		return te.Errors
	}
	return []string{err.Error()}
}
