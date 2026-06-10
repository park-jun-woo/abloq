//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what signingKey가 AWS 문서의 파생 키 벡터와 일치하는지 검증 (20150830/us-east-1/iam)
package cflog

import (
	"encoding/hex"
	"testing"
)

func TestSigningKey(t *testing.T) {
	got := hex.EncodeToString(signingKey("wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY", "20150830", "us-east-1", "iam"))
	want := "c4afb1cc5771d871763a393e44b703571b55cc28424d1a5e86da6ed3c154a4b9"
	if got != want {
		t.Errorf("signingKey = %s, want %s", got, want)
	}
}
