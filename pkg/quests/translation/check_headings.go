//ff:func feature=quest type=rule control=sequence topic=lossless
//ff:what 패리티 ① — 전체 헤딩 레벨 시퀀스(자유 헤딩 포함)와 인식 섹션 키 시퀀스(수·순서)의 원문↔번역 일치 검사
package translation

import (
	"fmt"

	"github.com/park-jun-woo/reins/pkg/quest"

	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

// checkHeadings compares the full heading-level sequence (free-form headings
// included) and the recognized section key sequence (count and order, per the
// blog.yaml heading maps) between origin and translation.
func checkHeadings(where string, o, t *agate.Doc) []quest.Fact {
	var facts []quest.Fact
	oLv, tLv := fmt.Sprint(headingLevels(o)), fmt.Sprint(headingLevels(t))
	if oLv != tLv {
		facts = append(facts, quest.Fact{Where: where + "#headings",
			Expected: "origin heading level sequence " + oLv, Actual: tLv})
	}
	oKeys, tKeys := fmt.Sprint(sectionKeys(o)), fmt.Sprint(sectionKeys(t))
	if oKeys != tKeys {
		facts = append(facts, quest.Fact{Where: where + "#sections",
			Expected: "origin recognized section sequence " + oKeys, Actual: tKeys})
	}
	return facts
}
