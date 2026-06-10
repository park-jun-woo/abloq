//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what [tag-taxonomy] geo.taxonomy 밖 태그 검출 — taxonomy 미선언(빈 목록)이면 조용히 스킵
//ff:why 보수적 스킵(Phase010 출처 헤딩 선례) — taxonomy는 선택적 SSOT 키라 부재가 위반이 아니다
package cluster

import (
	"slices"
	"strings"
)

// taxonomyViolation flags the article's tags that are missing from the
// declared geo.taxonomy vocabulary. An undeclared taxonomy skips the check.
func taxonomyViolation(tags, taxonomy []string) *Violation {
	if len(taxonomy) == 0 {
		return nil
	}
	var bad []string
	for _, tag := range tags {
		if !slices.Contains(taxonomy, tag) {
			bad = append(bad, tag)
		}
	}
	if len(bad) == 0 {
		return nil
	}
	return &Violation{Rule: "tag-taxonomy", Detail: "tags not in geo.taxonomy: " + strings.Join(bad, ", ")}
}
