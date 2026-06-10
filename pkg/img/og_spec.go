//ff:type feature=image type=schema
//ff:what OG 이미지 생성 입력 — 제목/브랜드 라인/커스텀 폰트 경로/출력 경로 (1200×630 고정)
package img

// OGSpec declares one OG image. FontPath empty = embedded Go Bold (latin);
// CJK or RTL titles need a custom TTF via FontPath.
type OGSpec struct {
	Title    string
	Brand    string
	FontPath string
	Out      string
}
