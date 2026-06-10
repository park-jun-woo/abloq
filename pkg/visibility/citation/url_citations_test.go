//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what urlCitations가 url_citation 타입·비어있지 않은 URL만 수집하고 다른 어노테이션은 무시하는지 검증
package citation

import (
	"reflect"
	"testing"
)

func TestURLCitations(t *testing.T) {
	contents := []oaiContent{
		{Type: "output_text", Annotations: []oaiAnnotation{
			{Type: "url_citation", URL: "https://blog.test/a/"},
			{Type: "file_citation", URL: "https://skip.example.org/"},
			{Type: "url_citation", URL: ""},
		}},
		{Type: "refusal"},
		{Type: "output_text", Annotations: []oaiAnnotation{
			{Type: "url_citation", URL: "https://blog.test/b/"},
		}},
	}
	got := urlCitations(contents)
	want := []string{"https://blog.test/a/", "https://blog.test/b/"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("urlCitations = %v, want %v", got, want)
	}
	if got := urlCitations(nil); got != nil {
		t.Errorf("nil contents = %v, want nil", got)
	}
}
