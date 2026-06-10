//ff:func feature=archive type=client control=sequence
//ff:what 외부 응답을 영수증 response JSON으로 환원 — {"status_code", "body"}, JSON 본문은 원형 임베드·그 외는 문자열 (2KB 상한)
package archive

import (
	"bytes"
	"encoding/json"
)

// wrapResponse reduces an external HTTP exchange to receipt evidence.
// statusCode 0 means a transport error and body carries the error text.
func wrapResponse(statusCode int, body []byte) json.RawMessage {
	if len(body) > 2048 {
		body = body[:2048]
	}
	payload := map[string]any{"status_code": statusCode}
	if json.Valid(body) && len(bytes.TrimSpace(body)) > 0 {
		payload["body"] = json.RawMessage(body)
	} else {
		payload["body"] = string(body)
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return json.RawMessage(`{}`)
	}
	return data
}
