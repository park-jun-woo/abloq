//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what GSC API 공용 POST — Bearer 토큰 JSON 호출, 비2xx는 에러, 본문 1MB 상한
package gsc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// gscClient bounds every Search Console call.
var gscClient = &http.Client{Timeout: 60 * time.Second}

// postJSON executes one authorized JSON POST against the Search Console API
// and returns the response body. A non-2xx status is an error: unlike the
// archiver there is no receipt to absorb it — the caller aborts the run.
func postJSON(endpoint, token string, payload any) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := gscClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gsc api returned %d: %s", resp.StatusCode, data)
	}
	return data, nil
}
