//ff:func feature=archive type=client control=sequence
//ff:what SPN2 저장 1건 — POST {WAYBACK_BASE_URL}/save에 url 폼 제출, 2xx면 done·그 외 failed (응답 기록)
package archive

import (
	"net/http"
	"net/url"
)

// saveWayback submits one Save-Page-Now capture. The Wayback timestamp is
// the original-author-time evidence, so the response (job id) is kept in
// the receipt verbatim.
func saveWayback(p Pending) Item {
	endpoint := envOr("WAYBACK_BASE_URL", "https://web.archive.org") + "/save"
	form := url.Values{"url": {p.Target}, "capture_all": {"1"}}
	header := http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}
	if auth := waybackAuth(); auth != "" {
		header.Set("Authorization", auth)
	}
	item := Item{
		DeployID: p.DeployID,
		Kind:     p.Kind,
		Target:   p.Target,
		Request:  requestJSON(endpoint, p.Target),
		Status:   StatusFailed,
	}
	code, body, err := httpPost(endpoint, header, []byte(form.Encode()))
	if err != nil {
		item.Response = wrapResponse(0, []byte(err.Error()))
		return item
	}
	item.Response = wrapResponse(code, body)
	if code >= 200 && code < 300 {
		item.Status = StatusDone
	}
	return item
}
