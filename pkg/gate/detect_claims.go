//ff:func feature=gate type=parser control=iteration dimension=1 topic=evidence
//ff:what 공개 DetectClaims API — 글 1편의 수치 주장 전부를 출처 보유 여부와 함께 검출 (Phase010 스캐너 재사용)
package gate

// DetectClaims finds every numeric claim in the article body, with its
// paragraph-level sourcing state. The numeric-claim-sourced gate rule and the
// claim-source scanner (operational backend) share this detector.
func DetectClaims(d *Doc) []Claim {
	var out []Claim
	for _, p := range claimParas(d) {
		out = append(out, paraClaims(d, p)...)
	}
	return out
}
