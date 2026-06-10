//ff:func feature=gate type=parser control=sequence
//ff:what parseDoc이 픽스처 글에서 BodyStart/메인 이미지/저작자 표기/섹션 헤딩을 정확히 채우는지 검증
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDoc(t *testing.T) {
	b := loadGateBlog(t)
	data, err := os.ReadFile(filepath.Join("testdata", "articles", "pass.md"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	d := parseDoc(buildHeadingIndex(b), "en", string(data))
	if !d.HasFM {
		t.Fatal("want HasFM true")
	}
	if d.BodyStart != 7 {
		t.Errorf("BodyStart = %d, want 7", d.BodyStart)
	}
	if !d.FirstIsImage {
		t.Error("want FirstIsImage true")
	}
	if d.AttribLine != d.FirstContentLine+1 {
		t.Errorf("AttribLine = %d, want %d", d.AttribLine, d.FirstContentLine+1)
	}
	if len(d.Sections) != 4 {
		t.Fatalf("want 4 sections, got %v", d.Sections)
	}
	if d.Sections[0].Key != "related" || d.Sections[3].Key != "changelog" {
		t.Errorf("section keys = %v", d.Sections)
	}
	if len(d.BadLevel) != 0 {
		t.Errorf("want no BadLevel, got %v", d.BadLevel)
	}
}
