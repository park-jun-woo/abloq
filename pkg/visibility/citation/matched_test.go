//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what Matched가 도메인 부분문자열 매칭으로 자기 URL만 선별하고 빈 도메인이면 nil인지 검증
package citation

import (
	"reflect"
	"testing"
)

func TestMatched(t *testing.T) {
	urls := []string{
		"https://blog.test/tech/post-a/",
		"https://other.example.org/x",
		"https://blog.test/tech/post-b/",
		"",
	}
	got := Matched("blog.test", urls)
	want := []string{"https://blog.test/tech/post-a/", "https://blog.test/tech/post-b/"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Matched = %v, want %v", got, want)
	}
	if got := Matched("", urls); got != nil {
		t.Errorf("empty domain = %v, want nil", got)
	}
	if got := Matched("blog.test", nil); got != nil {
		t.Errorf("no urls = %v, want nil", got)
	}
}
