//ff:func feature=insight type=parser control=sequence
//ff:what fold 검증 — NFD 한글이 NFC와 합치, 대문자·ß 케이스 폴딩, 멱등
package insight

import "testing"

func TestFold(t *testing.T) {
	nfd := "\u1112\u1161\u11ab\u1100\u1173\u11af" // decomposed jamo (NFD)
	nfc := "\ud55c\uae00"                         // precomposed syllables (NFC)
	if got := fold(nfd); got != nfc {
		t.Errorf("want NFD input folded to NFC %q, got %q", nfc, got)
	}
	if fold("STRASSE") != fold("Straße") {
		t.Errorf("want case folding to unify STRASSE and Strasse-with-eszett, got %q vs %q", fold("STRASSE"), fold("Straße"))
	}
	if got := fold(fold("MiXeD " + nfd)); got != fold("MiXeD "+nfd) {
		t.Errorf("want fold idempotent, got %q", got)
	}
}
