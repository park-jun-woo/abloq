//ff:func feature=image type=parser control=sequence
//ff:what statResult가 전/후 파일 크기를 채우고 없는 파일에 에러를 내는지 검증
package img

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStatResult(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "a")
	dst := filepath.Join(dir, "b")
	if err := os.WriteFile(src, []byte("12345"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := os.WriteFile(dst, []byte("12"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	res, err := statResult(src, dst)
	if err != nil || res.SrcBytes != 5 || res.DstBytes != 2 || res.Dst != dst {
		t.Errorf("statResult = %+v, %v; want sizes 5/2", res, err)
	}
	if _, err := statResult(filepath.Join(dir, "missing"), dst); err == nil {
		t.Error("statResult(missing src) expected error, got nil")
	}
	if _, err := statResult(src, filepath.Join(dir, "missing")); err == nil {
		t.Error("statResult(missing dst) expected error, got nil")
	}
}
