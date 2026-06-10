//ff:type feature=gate type=schema
//ff:what 게이트 룰 1개 — 룰ID/설명/검사 함수, 레지스트리(Rules)와 실행기(Run)의 단위
package gate

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// Rule is one structure-gate rule. Check inspects the whole target and returns
// one diagnostic per violation (empty means the rule passes).
type Rule struct {
	ID    string
	Desc  string
	Check func(t *Target) []blogyaml.Diagnostic
}
