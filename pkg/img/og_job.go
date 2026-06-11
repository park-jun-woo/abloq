//ff:type feature=image type=schema
//ff:what AI OG 실행 단위 1건 — 안(variant) × 샘플 번호 × 출력 경로, 안×count 전개의 평탄화 산물
//ff:why 안×count 이중 루프를 잡 목록으로 평탄화 — 실행 루프(OGAI)가 단일 차원으로 남아 실패 집계가 단순해진다
package img

// ogJob is one flattened generation unit: which variant, which sample number
// and where the WebP goes.
type ogJob struct {
	variant OGVariant
	n       int
	out     string
}
