//ff:func feature=archive type=client control=sequence
//ff:what IndexNow 그룹 처리 — 키 인자(빈 값은 INDEXNOW_KEY env fallback)로 전체 target을 일괄 제출, 배치 응답을 영수증들로 전개. 키 부재·페이로드 실패는 전 항목 failed
package archive

import (
	"encoding/json"
	"net/http"
	"os"
)

// processIndexNow submits every target in one batch POST. The caller-given
// key (site row value) wins; an empty key falls back to the global
// INDEXNOW_KEY environment. IndexNow is a batch protocol, so the receipts of
// the batch share one response; a non-2xx batch response fails every target
// (retry rearms them together).
func processIndexNow(key string, pending []Pending) []Item {
	if len(pending) == 0 {
		return nil
	}
	endpoint := envOr("INDEXNOW_ENDPOINT", "https://api.indexnow.org/indexnow")
	if key == "" {
		key = os.Getenv("INDEXNOW_KEY")
	}
	if key == "" {
		return fanoutItems(pending, endpoint, wrapResponse(0, []byte("INDEXNOW_KEY is not set")), StatusFailed)
	}
	payload, err := indexNowPayload(key, pending)
	if err != nil {
		return fanoutItems(pending, endpoint, wrapResponse(0, []byte(err.Error())), StatusFailed)
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fanoutItems(pending, endpoint, wrapResponse(0, []byte(err.Error())), StatusFailed)
	}
	header := http.Header{
		"Content-Type": {"application/json; charset=utf-8"},
		"Accept":       {"application/json"},
	}
	code, respBody, err := httpPost(endpoint, header, body)
	if err != nil {
		return fanoutItems(pending, endpoint, wrapResponse(0, []byte(err.Error())), StatusFailed)
	}
	status := StatusFailed
	if code >= 200 && code < 300 {
		status = StatusDone
	}
	return fanoutItems(pending, endpoint, wrapResponse(code, respBody), status)
}
