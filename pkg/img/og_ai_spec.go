//ff:type feature=image type=schema
//ff:what AI OG 실행 입력 — slug/오버레이 텍스트(제목·브랜드·폰트)/직행·드래프트 출력 경로/다중 안 여부/안당 샘플 수
package img

// OGAISpec declares one AI OG run. The prompt and provider live per variant
// (OGVariant) — this spec carries what every variant shares. Multi=false
// writes the single candidate straight to OutDir/{Slug}.webp (the adopted
// path); Multi=true writes draft candidates to DraftDir/{Slug}/{variant}-{n}.webp
// for review and manual adoption. Count is samples per variant (min 1).
type OGAISpec struct {
	Slug     string
	Title    string
	Brand    string
	FontPath string
	OutDir   string
	DraftDir string
	Multi    bool
	Count    int
}
