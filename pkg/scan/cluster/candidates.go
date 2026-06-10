//ff:func feature=scan type=generator control=iteration dimension=1 topic=cluster
//ff:what 위반 글 1편의 연결 후보 산출 — 자기 자신·완전 연결 쌍 제외, 정렬 후 상위 5건
package cluster

import "sort"

// maxCandidates caps the suggestions per violating article; a smaller corpus
// yields as many as it has.
const maxCandidates = 5

// candidates ranks every other article as a link suggestion for one
// violating article and keeps the top maxCandidates. Self references and
// already fully connected pairs never qualify.
func candidates(p post, posts []post, edges map[string]bool) []Candidate {
	self := PostKey(p.Section, p.Slug)
	pool := make([]ranked, 0, len(posts))
	for _, o := range posts {
		key := PostKey(o.Section, o.Slug)
		if key == self {
			continue
		}
		dirs := directions(p, o, edges)
		if len(dirs) == 0 {
			continue
		}
		pool = append(pool, ranked{
			cand:        Candidate{Section: o.Section, Slug: o.Slug, SharedTags: sharedTags(p.Tags, o.Tags), Directions: dirs},
			sameSection: o.Section == p.Section,
			dateDist:    dateDistance(p.Date, o.Date),
			key:         key,
		})
	}
	sort.Slice(pool, func(i, j int) bool { return lessRanked(pool[i], pool[j]) })
	if len(pool) > maxCandidates {
		pool = pool[:maxCandidates]
	}
	out := make([]Candidate, 0, len(pool))
	for _, r := range pool {
		out = append(out, r.cand)
	}
	return out
}
