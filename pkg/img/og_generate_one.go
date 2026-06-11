//ff:func feature=image type=generator control=sequence
//ff:what AI OG 1건 생성 — Provider 호출 → 1200×630 센터 크롭 → 흰 배경 평탄화 → (옵션) 텍스트 오버레이 → WebP 기록
package img

import "context"

// ogGenerateOne runs one job through the post-processing pipeline: the
// provider's raw image is normalized to the OG format, optionally overlaid
// with the deterministic title/brand composition, and written as WebP.
func ogGenerateOne(ctx context.Context, spec OGAISpec, job ogJob) error {
	const w, h = 1200, 630
	raw, err := job.variant.Provider.Generate(ctx, job.variant.Prompt)
	if err != nil {
		return err
	}
	canvas := FlattenWhite(CropCenter(raw, w, h))
	if job.variant.Overlay {
		text := OGSpec{Title: spec.Title, Brand: spec.Brand, FontPath: spec.FontPath}
		if err := OverlayText(canvas, text); err != nil {
			return err
		}
	}
	return SaveWebP(job.out, canvas)
}
