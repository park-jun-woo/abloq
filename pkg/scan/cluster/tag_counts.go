//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 태그별 보유 글 수 집계 — no-orphan-tag 판정(1편 보유 = 고아)의 입력
package cluster

// tagCounts counts how many articles carry each tag across the
// default-language corpus.
func tagCounts(posts []post) map[string]int64 {
	counts := map[string]int64{}
	for _, p := range posts {
		for _, tag := range p.Tags {
			counts[tag]++
		}
	}
	return counts
}
