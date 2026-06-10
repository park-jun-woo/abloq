//ff:func feature=archive type=client control=sequence
//ff:what IndexNow 일괄 페이로드 조립 — {host, key, keyLocation, urlList}, host는 첫 target에서·keyLocation은 프로토콜 규약 위치 (키 검증 자료 포함)
package archive

import (
	"fmt"
	"net/url"
)

// indexNowPayload builds the batch submission body. keyLocation defaults to
// the protocol-mandated https://<host>/<key>.txt and can be overridden with
// INDEXNOW_KEY_LOCATION — the endpoint verifies key ownership there.
func indexNowPayload(key string, pending []Pending) (map[string]any, error) {
	u, err := url.Parse(pending[0].Target)
	if err != nil {
		return nil, err
	}
	if u.Host == "" {
		return nil, fmt.Errorf("indexnow: target %q has no host", pending[0].Target)
	}
	keyLocation := envOr("INDEXNOW_KEY_LOCATION", "https://"+u.Host+"/"+key+".txt")
	return map[string]any{
		"host":        u.Host,
		"key":         key,
		"keyLocation": keyLocation,
		"urlList":     targetList(pending),
	}, nil
}
