//ff:func feature=sitesyaml type=parser control=sequence
//ff:what joinPath가 빈 prefix면 key만, 아니면 "prefix.key"를 반환하는지 검증
package sitesyaml

import "testing"

func TestJoinPath(t *testing.T) {
	if got := joinPath("", "sites"); got != "sites" {
		t.Errorf(`joinPath("", "sites") = %q, want "sites"`, got)
	}
	if got := joinPath("sites[0]", "name"); got != "sites[0].name" {
		t.Errorf(`joinPath("sites[0]", "name") = %q, want "sites[0].name"`, got)
	}
}
