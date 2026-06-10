//ff:func feature=claudemd type=generator control=sequence
//ff:what 테스트 픽스처 — 2개 언어/2개 섹션/sources 헤딩과 기본 임계값을 가진 Blog 제공
package claudemd

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

func testBlog() *blogyaml.Blog {
	return &blogyaml.Blog{
		Site:      blogyaml.Site{BaseURL: "https://t.example.com", Title: "T Blog", Author: "Tester"},
		Languages: []string{"ko", "en"},
		Sections:  []string{"opinion", "tech"},
		Structure: blogyaml.Structure{
			Order:    []string{"body", "sources"},
			Headings: map[string]map[string]string{"sources": {"ko": "출처", "en": "Sources"}},
		},
		Geo: blogyaml.Geo{FreshnessDays: 90, MinSources: 1, MinInternalLinks: 2, MinMeaningfulDiff: 10},
	}
}
