//ff:func feature=archive type=client control=sequence
//ff:what parsePrivateKey가 PKCS#8·PKCS#1 PEM을 받아들이고 비PEM·비RSA 키를 거부하는지 검증
package archive

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestParsePrivateKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}

	der8, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("pkcs8: %v", err)
	}
	pem8 := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der8})
	if _, err := parsePrivateKey(string(pem8)); err != nil {
		t.Errorf("PKCS#8: %v", err)
	}

	pem1 := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if _, err := parsePrivateKey(string(pem1)); err != nil {
		t.Errorf("PKCS#1: %v", err)
	}

	if _, err := parsePrivateKey("not pem at all"); err == nil {
		t.Error("non-PEM must fail")
	}

	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("ecdsa generate: %v", err)
	}
	derEC, err := x509.MarshalPKCS8PrivateKey(ecKey)
	if err != nil {
		t.Fatalf("ecdsa pkcs8: %v", err)
	}
	pemEC := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: derEC})
	if _, err := parsePrivateKey(string(pemEC)); err == nil {
		t.Error("non-RSA key must fail")
	}
}
