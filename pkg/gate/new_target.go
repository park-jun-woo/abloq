//ff:func feature=gate type=frame control=sequence
//ff:what 게이트 Target 조립 — 저장소 경로/blog.yaml/대상 글 목록을 받고 헤딩 인덱스를 파생해 채움
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// NewTarget assembles a gate target. arts usually comes from Discover, but a
// caller (e.g. reins) may pass any subset of articles to gate.
func NewTarget(dir string, b *blogyaml.Blog, arts []*Article) *Target {
	return &Target{Dir: dir, Blog: b, Articles: arts, heads: buildHeadingIndex(b)}
}
