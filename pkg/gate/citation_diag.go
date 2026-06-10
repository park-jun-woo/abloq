//ff:func feature=gate type=rule control=sequence topic=evidence
//ff:what 인용 1건의 진단 — 24h 내 영수증은 재검증 생략, retry 판정은 캐시하지 않고 RETRY 진단으로 보고
package gate

import (
	"net/http"
	"time"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// citationDiag judges one new citation. A fresh receipt skips re-verification;
// verified verdicts (except transient "retry") are written back to rcpts.
// ok is false when the citation passes (no diagnostic).
func citationDiag(file string, c Citation, rcpts map[string]receipt, client *http.Client, now time.Time) (blogyaml.Diagnostic, bool) {
	r, cached := rcpts[c.URL]
	if !cached || now.Sub(r.CheckedAt) >= citationTTL {
		verdict, detail := verifyCitation(client, c)
		r = receipt{CheckedAt: now, Verdict: verdict, Detail: detail}
		if verdict != "retry" {
			rcpts[c.URL] = r
		}
	}
	if r.Verdict == "ok" {
		return blogyaml.Diagnostic{}, false
	}
	msg := "citation URL " + c.URL + " is not reachable (" + r.Detail + ")"
	if r.Verdict == "meta-mismatch" {
		msg = "citation URL " + c.URL + ": " + r.Detail
	}
	if r.Verdict == "retry" {
		msg = "RETRY: citation URL " + c.URL + " hit a temporary failure (" + r.Detail + ") — re-run the gate"
	}
	return blogyaml.Diagnostic{File: file, Line: c.Line, Rule: "citation-exists", Message: msg}, true
}
