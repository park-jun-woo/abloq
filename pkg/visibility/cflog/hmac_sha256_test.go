//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what hmacSHA256·sha256Hex·sha256HexMAC이 알려진 다이제스트와 일치하는지 검증 (RFC 4231 케이스 2, 빈 문자열 SHA-256)
package cflog

import (
	"encoding/hex"
	"testing"
)

func TestHmacSHA256(t *testing.T) {
	// RFC 4231 test case 2: key "Jefe", data "what do ya want for nothing?"
	want := "5bdcc146bf60754e6a042426089575c75a003f089d2739839dec58b964ec3843"
	if got := hex.EncodeToString(hmacSHA256([]byte("Jefe"), "what do ya want for nothing?")); got != want {
		t.Errorf("hmacSHA256 = %s, want %s", got, want)
	}
	if got := sha256HexMAC([]byte("Jefe"), "what do ya want for nothing?"); got != want {
		t.Errorf("sha256HexMAC = %s, want %s", got, want)
	}
	if got := sha256Hex(nil); got != emptyPayloadHash {
		t.Errorf("sha256Hex(empty) = %s, want %s", got, emptyPayloadHash)
	}
}
