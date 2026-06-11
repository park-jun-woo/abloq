//ff:func feature=quest type=rule control=iteration dimension=1
//ff:what match 미출현 claim 중 REVIEW 기록에 disposition 라인이 없는 ID 목록 — review-record 커버리지 검사의 본체
package writing

import "github.com/park-jun-woo/abloq/pkg/insight"

// undisposed returns the IDs of match-missing claims the review record does
// not dispose of. Empty means full coverage.
func undisposed(review string, missing []insight.Claim) []string {
	var ids []string
	for _, c := range missing {
		if hasDisposition(review, c.ID) {
			continue
		}
		ids = append(ids, c.ID)
	}
	return ids
}
