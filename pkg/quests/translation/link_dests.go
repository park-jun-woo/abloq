//ff:func feature=quest type=parser control=iteration dimension=1 topic=lossless
//ff:what 본문 prose의 인라인 링크 목적지 전수 추출 — 라인별 lineDests 결합, 이미지와 코드 펜스 안 제외 (패리티 ⑤⑥ 공통)
package translation

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// linkDests extracts every inline (non-image) markdown link destination in
// the body prose. The external (⑤) and internal (⑥) parity checks filter it.
func linkDests(d *agate.Doc) []string {
	var dests []string
	for _, raw := range proseLines(d) {
		dests = append(dests, lineDests(raw)...)
	}
	return dests
}
