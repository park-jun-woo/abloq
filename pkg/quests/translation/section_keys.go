//ff:func feature=quest type=parser control=iteration dimension=1
//ff:what 인식된 섹션 헤딩 키 시퀀스 추출 — Doc.Sections(언어별 headings 맵으로 인식된 ## 헤딩)의 문서 순서 키 목록 (패리티 ①)
package translation

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// sectionKeys returns the document-order keys of the recognized section
// headings. Recognition is per-language (blog.yaml structure.headings), so an
// origin and its translation yield comparable key sequences.
func sectionKeys(d *agate.Doc) []string {
	var keys []string
	for _, s := range d.Sections {
		keys = append(keys, s.Key)
	}
	return keys
}
