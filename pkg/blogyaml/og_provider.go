//ff:func feature=blogyaml type=schema control=sequence
//ff:what image.og의 실효 provider 조회 — 미선언(빈 값) Blog도 local로 취급해 CLI 해석(플래그 > blog.yaml > local)이 안전하게 분기
package blogyaml

// OGProvider returns the effective OG provider: "local" (the default,
// including zero-value Blogs without an image block) or the declared name.
func (i Image) OGProvider() string {
	if i.OG.Provider == "" {
		return "local"
	}
	return i.OG.Provider
}
