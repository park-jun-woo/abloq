//ff:func feature=archive type=client control=sequence
//ff:what 테스트용 SA JSON 픽스처 생성 — 일회용 RSA 키를 PKCS#8 PEM으로 감싼 client_email/private_key JSON
package archive

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"testing"
)

// saJSONFixture builds a throwaway service-account JSON with a freshly
// generated RSA key — no real credential ever enters the repository.
func saJSONFixture(t *testing.T) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate rsa key: %v", err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("marshal pkcs8: %v", err)
	}
	pemText := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
	data, err := json.Marshal(map[string]string{
		"client_email": "stub@test-project.iam.gserviceaccount.com",
		"private_key":  string(pemText),
	})
	if err != nil {
		t.Fatalf("marshal sa json: %v", err)
	}
	return string(data)
}
