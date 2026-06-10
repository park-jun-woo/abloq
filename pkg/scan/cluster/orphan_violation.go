//ff:func feature=scan type=rule control=iteration dimension=1 topic=cluster
//ff:what [no-orphan-tag] 글 1편에만 붙은 태그 검출 — 코퍼스 전체 태그 카운트가 1인 보유 태그를 나열
package cluster

import "strings"

// orphanViolation flags the article's tags that no other article shares —
// an orphan tag clusters nothing.
func orphanViolation(tags []string, counts map[string]int64) *Violation {
	var orphans []string
	for _, tag := range tags {
		if counts[tag] == 1 {
			orphans = append(orphans, tag)
		}
	}
	if len(orphans) == 0 {
		return nil
	}
	return &Violation{Rule: "no-orphan-tag", Detail: "tags used by this article only: " + strings.Join(orphans, ", ")}
}
