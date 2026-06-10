//ff:func feature=archive type=client control=sequence
//ff:what SA private_key PEM 파싱 — PKCS#8 우선, PKCS#1 폴백, RSA 외 키는 거부
package archive

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// parsePrivateKey decodes the service-account PEM into an RSA signing key.
func parsePrivateKey(pemText string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemText))
	if block == nil {
		return nil, errors.New("private_key is not PEM")
	}
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		rsaKey, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("private_key is not an RSA key")
		}
		return rsaKey, nil
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
