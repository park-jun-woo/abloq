//ff:func feature=archive type=client control=sequence
//ff:what 공용 POST 실행 — 60초 타임아웃 클라이언트로 상태코드와 본문(1MB 상한)을 반환
package archive

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// archiveClient is the shared HTTP client: Wayback SPN2 captures can take
// tens of seconds, so the timeout is generous but bounded.
var archiveClient = &http.Client{Timeout: 60 * time.Second}

// httpPost executes one POST and returns the status code and response body.
// A non-2xx status is not an error — callers turn it into a failed receipt.
func httpPost(endpoint string, header http.Header, body []byte) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return 0, nil, err
	}
	if header != nil {
		req.Header = header
	}
	resp, err := archiveClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, data, nil
}
