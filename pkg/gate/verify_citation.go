//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what 인용 URL 1건 검증 — GET(리다이렉트 추적) 후 ok/broken/meta-mismatch/retry 판정, 5xx·네트워크 오류는 retry
package gate

import (
	"fmt"
	"io"
	"net/http"
)

// verifyCitation checks one citation URL: it must answer HTTP 200 (redirects
// followed) and its title/og:title must overlap the citation label. Network
// errors (incl. timeouts) and 5xx are transient — verdict "retry", not a hard
// failure.
func verifyCitation(client *http.Client, c Citation) (verdict, detail string) {
	resp, err := client.Get(c.URL)
	if err != nil {
		return "retry", trunc(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 500 {
		return "retry", fmt.Sprintf("HTTP %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return "broken", fmt.Sprintf("HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "retry", "read: " + trunc(err.Error())
	}
	if !metaMatch(string(body), c.Label) {
		return "meta-mismatch", "title/og:title does not overlap citation label " + trunc(c.Label)
	}
	return "ok", ""
}
