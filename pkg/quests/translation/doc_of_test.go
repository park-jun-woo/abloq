//ff:func feature=quest type=parser control=sequence
//ff:what 테스트 헬퍼 — 마크다운 문자열 1개를 인스턴스 blog.yaml 기준 지정 언어 Doc으로 파싱 (헤딩 인식이 언어별이라 lang 필수)
package translation

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	agate "github.com/park-jun-woo/abloq/pkg/gate"
)

func docOf(t *testing.T, lang, md string) *agate.Doc {
	t.Helper()
	b, diags, err := blogyaml.Load(writeInstance(t) + "/blog.yaml")
	if err != nil || len(diags) > 0 {
		t.Fatalf("fixture blog.yaml: err=%v diags=%v", err, diags)
	}
	return agate.ParseArticle(b, lang, md)
}
