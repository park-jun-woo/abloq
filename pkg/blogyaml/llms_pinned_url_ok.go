//ff:func feature=blogyaml type=rule control=sequence
//ff:what pinned url 형식 판정 — 절대 http(s) URL(호스트 필수) 또는 / 시작 경로만 참
package blogyaml

import (
	"net/url"
	"strings"
)

// llmsPinnedURLOK reports whether a pinned url is an absolute http(s) URL
// with a host, or a "/"-rooted site path.
func llmsPinnedURLOK(u string) bool {
	if strings.HasPrefix(u, "/") {
		return true
	}
	parsed, err := url.Parse(u)
	if err != nil {
		return false
	}
	return (parsed.Scheme == "http" || parsed.Scheme == "https") && parsed.Host != ""
}
