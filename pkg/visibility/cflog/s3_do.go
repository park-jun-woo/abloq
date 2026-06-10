//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 서명된 S3 GET 1회 실행 — sigv4 헤더를 붙여 보내고 비-2xx면 본문 머리를 담은 에러
package cflog

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// s3Do signs one GET request with SigV4 and executes it. Non-2xx responses
// are drained into an error carrying the body head (S3 XML error messages
// are short).
func (s S3Source) s3Do(rawURL string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	signRequest(req, s.AccessKey, s.SecretKey, s.SessionToken, s.Region, "s3", time.Now().UTC(), emptyPayloadHash)
	client := s.Client
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		head, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		resp.Body.Close()
		return nil, fmt.Errorf("s3: %s -> %s: %s", rawURL, resp.Status, head)
	}
	return resp, nil
}
