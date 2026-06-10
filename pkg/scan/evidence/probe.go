//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what HTTP 요청 1회 실행 — UA 헤더를 달아 상태 코드만 회수, 본문은 즉시 닫음 (리다이렉트는 클라이언트 기본 추적)
package evidence

import "net/http"

// probe performs one request and returns the final status code (redirects
// followed by the client). The body is discarded unread — the checker is a
// liveness probe, not a content fetch.
func (c *Checker) probe(method, target string) (int, error) {
	req, err := http.NewRequest(method, target, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.Client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}
