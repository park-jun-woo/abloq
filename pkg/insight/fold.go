//ff:func feature=insight type=parser control=sequence
//ff:what NFC 정규화 + 유니코드 케이스 폴딩 — 매칭 의미론의 정규화 단계 (형태소 분석 없음, 12언어·CJK 결정적)
package insight

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/unicode/norm"
)

// fold normalizes s to NFC then applies Unicode case folding. Both the body
// and every anchor pass through fold before substring matching; inflection
// and spelling variants are absorbed by the anchors synonym list instead.
func fold(s string) string {
	return cases.Fold().String(norm.NFC.String(s))
}
