//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what 글 1편의 신규 인용 목록을 검증해 진단 수집 — 통과 인용은 진단 없음
package gate

import (
	"net/http"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// citationDiags verifies one article's new citations against the receipt
// cache and collects the failing ones as diagnostics.
func citationDiags(file string, cs []Citation, rcpts map[string]receipt, client *http.Client) []blogyaml.Diagnostic {
	now := time.Now()
	var diags []blogyaml.Diagnostic
	for _, c := range cs {
		d, bad := citationDiag(file, c, rcpts, client, now)
		if bad {
			diags = append(diags, d)
		}
	}
	return diags
}
