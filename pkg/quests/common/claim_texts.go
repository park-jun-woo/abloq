//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what 파싱본의 수치 주장 텍스트 목록 추출 — exclude 해시 집합에 든 주장은 제외 (claim-scope 룰 전용)
package common

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// claimTexts lists the detected numeric-claim lines of one parsed article,
// skipping the claims whose text hash appears in exclude (the queue-payload
// authorization). A nil exclude keeps everything.
func claimTexts(d *agate.Doc, exclude map[string]bool) []string {
	var texts []string
	for _, c := range agate.DetectClaims(d) {
		if exclude[agate.HashText(c.Text)] {
			continue
		}
		texts = append(texts, c.Text)
	}
	return texts
}
