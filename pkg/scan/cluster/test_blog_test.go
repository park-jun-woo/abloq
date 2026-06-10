//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what 클러스터 스캐너 테스트 공통 Blog — ko 기본(루트 서빙)+en, tech 섹션, taxonomy 4종, min_internal_links 2
package cluster

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// testBlog builds the cluster test Blog in memory: ko is the root-served
// default language, en the translation, with the curated taxonomy the
// fixture posts violate in exactly one place (rogue).
func testBlog() *blogyaml.Blog {
	return &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://t.example.com", Title: "T", Author: "A", DefaultLangInSubdir: false},
		Languages: []string{"ko", "en"},
		Sections:  []string{"tech"},
		Geo: blogyaml.Geo{
			MinInternalLinks: 2,
			Taxonomy:         []string{"geo", "abloq", "hugo", "lonely"},
		},
	}
}
