//ff:func feature=archive type=client control=sequence
//ff:what SA 토큰 교환 — 사이트 자격(Keys)으로 SA를 로드해 GSC_TOKEN_URL에 jwt-bearer grant POST, scope 인자별 access_token 반환
//ff:why 토큰 발급도 외부 호출 — env로 스텁 지향 가능. scope를 인자로 받아 pkg/visibility/gsc가 같은 SA 자격을 재사용한다 (Phase008·013)
package archive

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// GSCTokenWith exchanges the service-account assertion for an access token
// of the given scope using the per-site credentials (empty Keys fall back to
// GSC_SA_JSON / GSC_SA_JSON_PATH). The token endpoint is env-overridable
// (GSC_TOKEN_URL) because the token issuance itself is an external call the
// Hurl stub must intercept.
func GSCTokenWith(keys Keys, scope string) (string, error) {
	sa, err := loadServiceAccount(keys.GSCSAJSON, keys.GSCSAJSONPath)
	if err != nil {
		return "", err
	}
	tokenURL := envOr("GSC_TOKEN_URL", "https://oauth2.googleapis.com/token")
	assertion, err := gscAssertion(sa, scope, tokenURL)
	if err != nil {
		return "", err
	}
	form := url.Values{
		"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"assertion":  {assertion},
	}
	header := http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json"},
	}
	code, body, err := httpPost(tokenURL, header, []byte(form.Encode()))
	if err != nil {
		return "", err
	}
	if code < 200 || code >= 300 {
		return "", fmt.Errorf("token endpoint returned %d", code)
	}
	var tok struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tok); err != nil {
		return "", err
	}
	if tok.AccessToken == "" {
		return "", errors.New("token endpoint returned no access_token")
	}
	return tok.AccessToken, nil
}
