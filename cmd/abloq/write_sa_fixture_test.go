//ff:func feature=cli type=command control=sequence
//ff:what 테스트용 GSC SA JSON 픽스처 — 일회용 RSA 키 생성 (실제 자격증명 무사용)
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"testing"
)

// testSAJSONFixture builds a throwaway service-account JSON for runArchive
// tests — a fresh RSA key per run, no real credential in the repository.
func testSAJSONFixture(t *testing.T) string {
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
