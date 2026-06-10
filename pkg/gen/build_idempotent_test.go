//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what Build 2회 연속 실행이 산출물 4종 전부에서 바이트 동일(no-op)인지 검증 — 멱등성 완료 판정
package gen

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestBuildIdempotent(t *testing.T) {
	dir := filepath.Join("testdata", "golden")
	b := loadGoldenBlog(t)
	first := Build(dir, b)
	second := Build(dir, b)
	for i, o := range first {
		if !bytes.Equal(o.Data, second[i].Data) {
			t.Errorf("%s differs between consecutive builds", o.Path)
		}
	}
}
