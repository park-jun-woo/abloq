//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what linkTargets가 본문 마크다운 링크 중 내부 글 키만 등장 순서·중복 제거로 수집하는지 검증
package cluster

import (
	"reflect"
	"testing"
)

func TestLinkTargets(t *testing.T) {
	body := "[a](/tech/thin/) 본문 [x](https://example.org/) [b](/tech/hub/)\n" +
		"중복 [a2](/tech/thin/) 그리고 [tag](/tags/geo/)\n"
	got := linkTargets(testBlog(), "ko", body)
	want := []string{"tech/thin", "tech/hub"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("linkTargets = %v, want %v", got, want)
	}
	if got := linkTargets(testBlog(), "ko", "링크 없음\n"); len(got) != 0 {
		t.Errorf("no links: got %v", got)
	}
}
