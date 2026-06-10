//ff:func feature=bots type=dict control=iteration dimension=1 topic=crawl
//ff:what candidateTokens가 비어 있지 않고 전부 소문자(부분일치 계약)인지 검증
package bots

import (
	"strings"
	"testing"
)

func TestCandidateTokens(t *testing.T) {
	toks := candidateTokens()
	if len(toks) == 0 {
		t.Fatal("empty token dictionary")
	}
	for _, tok := range toks {
		if tok == "" || tok != strings.ToLower(tok) {
			t.Errorf("token %q must be non-empty lowercase", tok)
		}
	}
}
