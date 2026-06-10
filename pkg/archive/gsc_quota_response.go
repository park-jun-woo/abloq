//ff:func feature=archive type=client control=sequence
//ff:what 429 쿼터 초과 응답을 영수증 evidence로 환원 — wrapResponse에 quota_exceeded: true 마커 추가
package archive

import "encoding/json"

// gscQuotaResponse marks a quota-exceeded 429 so operators (and a future
// scanner) can tell quota burn from real submission errors.
func gscQuotaResponse(statusCode int, body []byte) json.RawMessage {
	var payload map[string]any
	if err := json.Unmarshal(wrapResponse(statusCode, body), &payload); err != nil {
		return json.RawMessage(`{"status_code":429,"quota_exceeded":true}`)
	}
	payload["quota_exceeded"] = true
	data, err := json.Marshal(payload)
	if err != nil {
		return json.RawMessage(`{"status_code":429,"quota_exceeded":true}`)
	}
	return data
}
