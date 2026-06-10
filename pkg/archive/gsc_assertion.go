//ff:func feature=archive type=client control=sequence
//ff:what RS256 서명 JWT bearer assertion 조립 — iss=SA 이메일, scope는 인자(indexing/webmasters.readonly), aud=토큰 엔드포인트, 1시간 유효
//ff:why scope 하드코딩 금지: 스텁은 scope를 검증하지 않아 코드 판정은 통과한 채 본번에서만 403이 난다 — 호출자가 용도별 scope를 넘긴다 (Phase013)
package archive

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

// gscAssertion builds the signed JWT the OAuth2 jwt-bearer grant exchanges
// for a Google API access token of the given scope. Implemented with the
// standard library only — no external JWT dependency for one fixed-claim
// assertion.
func gscAssertion(sa *serviceAccount, scope, audience string) (string, error) {
	key, err := parsePrivateKey(sa.PrivateKey)
	if err != nil {
		return "", err
	}
	b64 := base64.RawURLEncoding
	header := b64.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	now := time.Now()
	claims, err := json.Marshal(map[string]any{
		"iss":   sa.ClientEmail,
		"scope": scope,
		"aud":   audience,
		"iat":   now.Unix(),
		"exp":   now.Add(time.Hour).Unix(),
	})
	if err != nil {
		return "", err
	}
	signing := header + "." + b64.EncodeToString(claims)
	digest := sha256.Sum256([]byte(signing))
	sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if err != nil {
		return "", err
	}
	return signing + "." + b64.EncodeToString(sig), nil
}
