//ff:func feature=quest type=rule control=sequence
//ff:what citation-exists 어댑터 발동 검증 — httptest 404 URL 인용에서 Fact 매핑 (Offline 아님 — 실 HTTP 경로)
package writing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdaptCitationExists(t *testing.T) {
	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()
	root := writeInstance(t)
	art, _ := passFixtures()
	cited := strings.Replace(art,
		"This body mentions the alpha anchor.",
		"This body mentions the alpha anchor ([Broken Ref]("+ts.URL+"/gone)).",
		1)
	fired, fact := fireRule(t, adaptRule("citation-exists"), subWith(t, root, cited, ""))
	if !fired {
		t.Fatal("citation-exists: want fired on an HTTP 404 citation")
	}
	if !strings.Contains(fact.Actual, ts.URL) {
		t.Errorf("Actual = %q, want the cited URL", fact.Actual)
	}
}
