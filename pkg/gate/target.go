//ff:type feature=gate type=schema
//ff:what 게이트 입력 1세트 — 저장소 경로 + blog.yaml + 대상 글 목록 + Offline 스위치 (+ 헤딩 인덱스), 모든 룰의 검사 대상
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Target is the unit a gate run inspects: the repository root (Dir, where
// blog.yaml lives), the parsed blog.yaml, and the articles under review.
// Offline skips the network-dependent rules (citation-exists).
type Target struct {
	Dir      string
	Blog     *blogyaml.Blog
	Articles []*Article
	Offline  bool

	heads headingIndex // derived from Blog.Structure (built by NewTarget)
}
