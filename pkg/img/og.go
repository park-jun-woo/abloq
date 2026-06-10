//ff:func feature=image type=generator control=sequence
//ff:what OG 생성 본체 — 1200×630 렌더 후 spec.Out에 WebP로 기록
package img

// OG renders spec and writes the WebP to spec.Out.
func OG(spec OGSpec) error {
	m, err := RenderOG(spec)
	if err != nil {
		return err
	}
	return SaveWebP(spec.Out, m)
}
