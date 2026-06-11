//ff:func feature=quest type=generator control=sequence
//ff:what headingMapLines 검증 — order 순서의 원문→대상 헤딩 라인 렌더, 헤딩 미선언 키(body) 생략
package translation

import (
	"fmt"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestHeadingMapLines(t *testing.T) {
	root := writeInstance(t)
	b, _, err := blogyaml.Load(root + "/blog.yaml")
	if err != nil {
		t.Fatalf("blog.yaml: %v", err)
	}
	got := fmt.Sprint(headingMapLines(b, "en", "ja"))
	want := `[- sources: "Sources" -> "出典"]`
	if got != want {
		t.Errorf("lines = %s, want %s", got, want)
	}
}
