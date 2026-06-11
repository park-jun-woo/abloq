//ff:func feature=quest type=rule control=iteration dimension=1 topic=lossless
//ff:what 패리티 ⑦ fm-mirror — 번역 date·lastmod == 원문 date·lastmod (ParseFMTime 비교), lastmod 위조 치즈 차단
//ff:why 이식 "규약"만으로는 lastmod 위조가 무검증(채택 룰 어느 것도 안 봄) — honest-lastmod는 git 기준선·큐 전제라 이 체인에서 배제하고 ⑦이 대체한다. 원문 값 부재는 Prepare 에러라 여기 도달 시 원문은 항상 파싱된다 (Phase017 계획 ⑦)
package translation

import (
	"github.com/park-jun-woo/reins/pkg/quest"
)

// checkFMMirror requires the translation's date and lastmod to mirror the
// origin's exactly (same parsed instants). The origin side is guaranteed
// parseable by Prepare; a missing or differing translation value fires.
func checkFMMirror(sub *Submission) []quest.Fact {
	var facts []quest.Fact
	for _, key := range []string{"date", "lastmod"} {
		oT, _ := fmTime(sub.Origin.Doc, key)
		tT, ok := fmTime(sub.Target.Articles[0].Doc, key)
		if ok && tT.Equal(oT) {
			continue
		}
		facts = append(facts, quest.Fact{Where: sub.Article + "#" + key,
			Expected: key + " mirrors the origin (" + oT.Format("2006-01-02T15:04:05Z07:00") + ")",
			Actual:   fmValue(sub.Target.Articles[0].Doc, key)})
	}
	return facts
}
