//ff:func feature=image type=client control=sequence
//ff:what Gemini API 공용 POST — x-goog-api-key 헤더 JSON 호출, 비2xx는 본문 동봉 에러, 본문 32MB 상한 (stdlib만)
package ogprovider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// geminiClient bounds every generation call — image models can take a while,
// but never unbounded.
var geminiClient = &http.Client{Timeout: 120 * time.Second}

// postJSON executes one JSON POST against the Gemini API and returns the
// response body. A non-2xx status is an error carrying a body snippet so
// quota/auth failures diagnose themselves.
func postJSON(ctx context.Context, endpoint, key string, payload any) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", key)
	resp, err := geminiClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		snippet := data
		if len(snippet) > 512 {
			snippet = snippet[:512]
		}
		return nil, fmt.Errorf("gemini api %s: %s", resp.Status, snippet)
	}
	return data, nil
}
