//ff:func feature=gate type=rule control=iteration dimension=1 topic=evidence
//ff:what [citation-exists] 신규 인용 URL의 실재 검증 — HTTP 200+메타 일치, 영수증 캐시, Offline이면 스킵
//ff:why 판정 등급 분리 — 게이트는 실재(URL 200+메타 일치, 결정적)까지만 자동 판정한다. 지지(출처가 주장을 실제로 뒷받침하는가)는 비결정이라 룰이 아니며, 퀘스트의 REVIEW 산출물(검토 기록 존재)로만 확인한다
package gate

import (
	"net/http"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// citationHTTP is the client citation-exists uses: redirects are followed
// (net/http default), slow hosts time out into a "retry" verdict.
var citationHTTP = &http.Client{Timeout: 10 * time.Second}

// ruleCitationExists verifies that every citation URL added since git HEAD
// exists: HTTP 200 after redirects, with a title/og:title that overlaps the
// citation label. Verdicts are cached as receipt files for 24h. The whole
// rule is skipped offline.
func ruleCitationExists(t *Target) []blogyaml.Diagnostic {
	if t.Offline {
		return nil
	}
	rcpts := loadReceipts(t.Dir)
	checked := false
	var diags []blogyaml.Diagnostic
	for _, a := range t.Articles {
		cs := newCitations(a)
		checked = checked || len(cs) > 0
		diags = append(diags, citationDiags(a.Path, cs, rcpts, citationHTTP)...)
	}
	if checked {
		_ = saveReceipts(t.Dir, rcpts) // a cache write failure must not fail the gate
	}
	return diags
}
