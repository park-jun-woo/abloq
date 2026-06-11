//ff:func feature=blogyaml type=schema control=sequence
//ff:what geo.llms_txt의 실효 mode 조회 — 미주입(영값) Blog도 auto로 취급해 생성 게이트·렌더러가 안전하게 분기
package blogyaml

// LlmsTxtMode returns the effective llms.txt mode: "auto" (the default,
// including zero-value Blogs built without defaultBlog), "manual" or "off".
func (g Geo) LlmsTxtMode() string {
	if g.LlmsTxt.Mode == "" {
		return "auto"
	}
	return g.LlmsTxt.Mode
}
