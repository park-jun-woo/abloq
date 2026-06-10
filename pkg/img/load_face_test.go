//ff:func feature=image type=parser control=sequence
//ff:what LoadFace가 빈 경로에서 임베디드 Go Bold를 로드하고 없는 경로에 에러를 내는지 검증
package img

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFace(t *testing.T) {
	face, err := LoadFace("", 32)
	if err != nil {
		t.Fatalf("LoadFace(embedded): %v", err)
	}
	if face == nil {
		t.Fatal("LoadFace returned nil face")
	}
	if _, err := LoadFace("/nonexistent/font.ttf", 32); err == nil {
		t.Error("LoadFace(missing path) expected error, got nil")
	}
	garbage := filepath.Join(t.TempDir(), "bad.ttf")
	if err := os.WriteFile(garbage, []byte("not a font"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := LoadFace(garbage, 32); err == nil {
		t.Error("LoadFace(garbage bytes) expected parse error, got nil")
	}
}
