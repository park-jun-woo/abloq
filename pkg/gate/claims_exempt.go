//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what 공개 ClaimsExempt API — 글 단위 claims_ignore 예외(사유 필수) 여부, Phase010 스캐너가 게이트와 같은 예외를 적용
package gate

// ClaimsExempt reports whether the article opts out of numeric-claim
// detection via a valid claims_ignore (a non-empty list of reason strings).
// A malformed claims_ignore does not exempt — the gate flags it, and the
// scanner keeps scanning. The claim-source scanner shares this judgement.
func ClaimsExempt(a *Article) bool {
	exempt, _ := claimsIgnore(a)
	return exempt
}
