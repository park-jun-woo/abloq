//ff:func feature=content type=schema control=sequence
//ff:what Entry JSON 직렬화에 summary가 새지 않는지 검증 — posts 업서트 계약(컬럼 1:1) 바이트 불변, Summary는 내부 전용(json:"-")
package content

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEntryJSONExcludesSummary(t *testing.T) {
	e := Entry{Lang: "en", Section: "tech", Slug: "x", Title: "T", Summary: "leak me"}
	data, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(data)
	if strings.Contains(s, "leak me") || strings.Contains(s, "summary") || strings.Contains(s, "Summary") {
		t.Errorf("Entry JSON must not carry summary (posts contract): %s", s)
	}
}
