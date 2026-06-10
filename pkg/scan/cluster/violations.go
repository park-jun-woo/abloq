//ff:func feature=scan type=rule control=sequence topic=cluster
//ff:what 글 1편의 클러스터 위반 4종을 고정 순서로 수집 — tag-taxonomy → no-orphan-tag → min-internal-links → no-isolated-post
//ff:why 순서가 payload JSON 바이트를 결정한다 — CLI·endpoint diff 등가와 멱등 적재의 전제
package cluster

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// violations runs the four cluster checks over one node, in the fixed rule
// order the payload serialization depends on.
func violations(p post, b *blogyaml.Blog, tags map[string]int64, inlinks map[string]int64) []Violation {
	found := make([]Violation, 0, 4)
	if v := taxonomyViolation(p.Tags, b.Geo.Taxonomy); v != nil {
		found = append(found, *v)
	}
	if v := orphanViolation(p.Tags, tags); v != nil {
		found = append(found, *v)
	}
	if v := linksViolation(len(p.Outlinks), b.Geo.MinInternalLinks); v != nil {
		found = append(found, *v)
	}
	if v := isolationViolation(inlinks[PostKey(p.Section, p.Slug)]); v != nil {
		found = append(found, *v)
	}
	return found
}
