//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what blog.yaml 선언 언어·섹션의 전 글 수집 + git HEAD 원본 부착 — 게이트 기본 대상 목록
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Discover collects every article under dir/content/{lang}/{section}/ for the
// declared languages and sections, with git HEAD baselines attached.
func Discover(dir string, b *blogyaml.Blog) []*Article {
	hi := buildHeadingIndex(b)
	var arts []*Article
	for _, lang := range b.Languages {
		arts = append(arts, discoverLang(dir, b, hi, lang)...)
	}
	attachBaselines(dir, hi, arts)
	return arts
}
