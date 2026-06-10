//ff:type feature=gate type=schema
//ff:what 글 1편의 구조 파싱본 — front matter 원문, 본문 라인, 첫 이미지/저작자 표기 위치, 인식된 섹션 헤딩 목록
package gate

// Doc is the parsed structural view of one article: the raw front matter, the
// body split into lines, and the structural features the gate rules inspect.
type Doc struct {
	FrontMatter string // raw text between the --- fences (fences excluded)
	HasFM       bool   // a well-formed front matter block was found
	Body        string // everything after the closing fence
	BodyLines   []string
	BodyStart   int // 1-based file line number of BodyLines[0]

	FirstContentLine int  // index in BodyLines of the first non-blank line (-1 if none)
	FirstIsImage     bool // the first non-blank body line is a markdown image
	AttribLine       int  // index of the attribution line right after the image (-1 if none)

	Sections []SectionHit // recognized level-2 section headings, in document order
	BadLevel []SectionHit // headings recognized by name but NOT at level 2
}
