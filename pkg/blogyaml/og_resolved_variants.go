//ff:func feature=blogyaml type=schema control=iteration dimension=1
//ff:what 선언된 전 안을 상속 병합해 확정 목록으로 반환 — --all-variants 경로의 입력, 선언 순서 보존
package blogyaml

// ResolvedVariants returns every declared variant merged over the site-wide
// defaults, in declaration order. An empty declaration yields nil — the
// caller falls back to DefaultVariant.
func (o ImageOG) ResolvedVariants() []OGVariantSpec {
	var out []OGVariantSpec
	base := o.DefaultVariant()
	for _, d := range o.Variants {
		out = append(out, ogResolveVariant(base, d))
	}
	return out
}
