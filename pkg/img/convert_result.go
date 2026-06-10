//ff:type feature=image type=schema
//ff:what 이미지 변환 결과 — 출력 경로와 전/후 바이트 수 (CLI 절감률 출력용)
package img

// ConvertResult reports one WebP conversion.
type ConvertResult struct {
	Dst      string
	SrcBytes int64
	DstBytes int64
}
