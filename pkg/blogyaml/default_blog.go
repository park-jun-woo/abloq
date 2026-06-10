//ff:func feature=blogyaml type=parser control=sequence
//ff:what 스키마 v1 기본값이 채워진 Blog를 생성 — strict 디코드가 이 위에 덮어써서 기본값 주입을 구현
//ff:why 포인터 필드 없이 "미지정 = 기본값, 명시 0 = 검증 대상"을 구분하기 위해 디코드 전 선주입 방식 채택
package blogyaml

// defaultBlog returns a Blog pre-filled with schema v1 defaults.
// Decoding blog.yaml over this value leaves absent keys at their defaults.
func defaultBlog() Blog {
	return Blog{
		Site: Site{
			DefaultLangInSubdir: true,
		},
		Geo: Geo{
			LlmsTxt:           "auto",
			JSONLD:            []string{"Article", "Person"},
			FreshnessDays:     90,
			MinSources:        1,
			MinInternalLinks:  2,
			MinMeaningfulDiff: 10,
			PriorityWeights: PriorityWeights{
				Fetcher:  3,
				Train:    1,
				GSC:      1,
				Citation: 2,
			},
		},
		Deploy: Deploy{
			Provider:  "s3-cloudfront",
			Terraform: false,
			IndexNow:  true,
		},
	}
}
