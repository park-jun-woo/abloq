//ff:func feature=insight type=rule control=sequence
//ff:what claim 출현 판정 검증 — 동의어 중 1개 출현이면 true, 전부 미출현·anchors 없음은 false, NFC+폴딩 매칭
package insight

import "testing"

func TestClaimFound(t *testing.T) {
	nfcBody := "The machine decides the rest. \ud55c\uae00." // NFC syllables
	nfdAnchor := "\u1112\u1161\u11ab\u1100\u1173\u11af"      // same word, NFD jamo
	body := fold(nfcBody)
	if !claimFound(body, Claim{Anchors: []string{"absent phrase", "MACHINE DECIDES"}}) {
		t.Errorf("want true when one synonym anchor folds into the body")
	}
	if !claimFound(body, Claim{Anchors: []string{nfdAnchor}}) {
		t.Errorf("want true for NFD anchor against NFC body")
	}
	if claimFound(body, Claim{Anchors: []string{"absent phrase"}}) {
		t.Errorf("want false when no anchor occurs")
	}
	if claimFound(body, Claim{}) {
		t.Errorf("want false for anchorless claim")
	}
}
