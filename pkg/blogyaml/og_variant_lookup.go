//ff:func feature=blogyaml type=schema control=iteration dimension=1
//ff:what 이름으로 안 1개 조회(상속 병합 완료 형태) — --variant 콤마 목록 해석의 입력, 미선언 이름은 false
package blogyaml

// Variant finds one declared variant by name and returns it merged over the
// site-wide defaults.
func (o ImageOG) Variant(name string) (OGVariantSpec, bool) {
	base := o.DefaultVariant()
	for _, d := range o.Variants {
		if d.Name == name {
			return ogResolveVariant(base, d), true
		}
	}
	return OGVariantSpec{}, false
}
