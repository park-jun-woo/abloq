//ff:func feature=image type=generator control=sequence
//ff:what 이미지 변환 본체 — 디코드 → 흰 배경 평탄화 → 가로 maxW 축소 → WebP 기록, 전/후 크기 반환
package img

// Convert turns src (png/jpeg/gif/webp) into a WebP at dst, flattened onto
// white and resized to at most maxW pixels wide (0 = no resize).
func Convert(src, dst string, maxW int) (ConvertResult, error) {
	m, err := DecodeAny(src)
	if err != nil {
		return ConvertResult{}, err
	}
	flat := ResizeMax(FlattenWhite(m), maxW)
	if err := SaveWebP(dst, flat); err != nil {
		return ConvertResult{}, err
	}
	return statResult(src, dst)
}
