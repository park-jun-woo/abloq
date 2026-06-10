//ff:func feature=archive type=client control=sequence
//ff:what gscAssertion이 header.claims.signature 3부 JWT를 만들고 클레임(iss/scope 인자/aud)과 RS256 서명이 유효한지 검증
package archive

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestGscAssertion(t *testing.T) {
	saJSON := saJSONFixture(t)
	var sa serviceAccount
	if err := json.Unmarshal([]byte(saJSON), &sa); err != nil {
		t.Fatalf("fixture: %v", err)
	}
	assertion, err := gscAssertion(&sa, ScopeWebmastersReadonly, "https://token.test/token")
	if err != nil {
		t.Fatalf("gscAssertion: %v", err)
	}
	parts := strings.Split(assertion, ".")
	if len(parts) != 3 {
		t.Fatalf("JWT has %d parts, want 3", len(parts))
	}
	claimsRaw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatalf("decode claims: %v", err)
	}
	var claims map[string]any
	if err := json.Unmarshal(claimsRaw, &claims); err != nil {
		t.Fatalf("unmarshal claims: %v", err)
	}
	if claims["iss"] != sa.ClientEmail || claims["aud"] != "https://token.test/token" ||
		claims["scope"] != "https://www.googleapis.com/auth/webmasters.readonly" {
		t.Errorf("claims = %v", claims)
	}

	key, err := parsePrivateKey(sa.PrivateKey)
	if err != nil {
		t.Fatalf("parse key: %v", err)
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		t.Fatalf("decode signature: %v", err)
	}
	digest := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if err := rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, digest[:], sig); err != nil {
		t.Errorf("RS256 signature invalid: %v", err)
	}

	if _, err := gscAssertion(&serviceAccount{ClientEmail: "x", PrivateKey: "bad"}, ScopeIndexing, "aud"); err == nil {
		t.Error("bad key must fail")
	}
}
