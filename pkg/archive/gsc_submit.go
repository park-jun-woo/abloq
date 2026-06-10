//ff:func feature=archive type=client control=selection
//ff:what GSC publish 1건 — 2xx done · 429(쿼터 초과) failed로 기록(quota_exceeded 마킹) · 그 외 failed
package archive

import (
	"encoding/json"
	"net/http"
)

// gscSubmit publishes one URL_UPDATED notification. A server-side 429 means
// the daily quota is already burnt: the receipt is recorded as failed with
// the 429 evidence and a quota_exceeded marker — retry rearms it for a
// later (post-reset) process run.
func gscSubmit(p Pending, endpoint, token string) Item {
	item := Item{
		DeployID: p.DeployID,
		Kind:     p.Kind,
		Target:   p.Target,
		Request:  requestJSON(endpoint, p.Target),
		Status:   StatusFailed,
	}
	body, err := json.Marshal(map[string]string{"url": p.Target, "type": "URL_UPDATED"})
	if err != nil {
		item.Response = wrapResponse(0, []byte(err.Error()))
		return item
	}
	header := http.Header{
		"Content-Type":  {"application/json"},
		"Accept":        {"application/json"},
		"Authorization": {"Bearer " + token},
	}
	code, respBody, err := httpPost(endpoint, header, body)
	switch {
	case err != nil:
		item.Response = wrapResponse(0, []byte(err.Error()))
	case code == http.StatusTooManyRequests:
		item.Response = gscQuotaResponse(code, respBody)
	case code >= 200 && code < 300:
		item.Status = StatusDone
		item.Response = wrapResponse(code, respBody)
	default:
		item.Response = wrapResponse(code, respBody)
	}
	return item
}
