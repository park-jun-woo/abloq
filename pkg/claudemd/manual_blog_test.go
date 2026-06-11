//ff:func feature=claudemd type=generator control=sequence
//ff:what 테스트 픽스처 — llms_txt mode manual인 Blog 제공 (생성물 목록 제외 검증용)
package claudemd

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

func manualBlog() *blogyaml.Blog {
	b := testBlog()
	b.Geo.LlmsTxt = blogyaml.LlmsTxtSpec{Mode: "manual", Languages: []string{"base"}}
	return b
}
