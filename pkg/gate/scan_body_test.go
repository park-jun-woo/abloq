//ff:func feature=gate type=parser control=sequence
//ff:what scanBody가 첫 콘텐츠 라인/이미지 판정과 레벨별 섹션 분류(##=Sections, 그 외=BadLevel)를 수행하는지 검증
package gate

import (
	"strings"
	"testing"
)

func TestScanBody(t *testing.T) {
	b := loadGateBlog(t)
	d := &Doc{FirstContentLine: -1, AttribLine: -1, BodyStart: 1}
	d.Body = "\n![img](/i.webp)\n\ntext\n\n## Sources\n\n### Related\n"
	d.BodyLines = strings.Split(d.Body, "\n")
	scanBody(buildHeadingIndex(b), "en", d)
	if d.FirstContentLine != 1 || !d.FirstIsImage {
		t.Errorf("first content = %d image=%v, want 1 true", d.FirstContentLine, d.FirstIsImage)
	}
	if len(d.Sections) != 1 || d.Sections[0].Key != "sources" || d.Sections[0].Line != 5 {
		t.Errorf("Sections = %v, want sources at line 5", d.Sections)
	}
	if len(d.BadLevel) != 1 || d.BadLevel[0].Key != "related" || d.BadLevel[0].Level != 3 {
		t.Errorf("BadLevel = %v, want 3-level related", d.BadLevel)
	}
}
