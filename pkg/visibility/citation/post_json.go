//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what 엔진 API 공용 POST — 헤더 지정 JSON 호출, 비2xx는 에러, 본문 4MB 상한 (stdlib만, 신규 의존성 0)
package citation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// engineClient bounds every sampling call — engine answers with web search
// can take a while, but never unbounded.
var engineClient = &http.Client{Timeout: 120 * time.Second}

// postJSON executes one JSON POST against an engine API and returns the
// response body. A non-2xx status is an error — the runner records it as
// the sample's evidence.
func postJSON(endpoint string, header http.Header, payload any) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := engineClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("engine api returned %d: %s", resp.StatusCode, data)
	}
	return data, nil
}
