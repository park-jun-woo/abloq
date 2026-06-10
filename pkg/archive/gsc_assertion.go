//ff:func feature=archive type=client control=sequence
//ff:what RS256 서명 JWT bearer assertion 조립 — iss=SA 이메일, scope=indexing, aud=토큰 엔드포인트, 1시간 유효
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
// for an Indexing API access token. Implemented with the standard library
// only — no external JWT dependency for one fixed-claim assertion.
func gscAssertion(sa *serviceAccount, audience string) (string, error) {
	key, err := parsePrivateKey(sa.PrivateKey)
	if err != nil {
		return "", err
	}
	b64 := base64.RawURLEncoding
	header := b64.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	now := time.Now()
	claims, err := json.Marshal(map[string]any{
		"iss":   sa.ClientEmail,
		"scope": "https://www.googleapis.com/auth/indexing",
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
