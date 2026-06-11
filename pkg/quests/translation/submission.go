//ff:type feature=quest type=schema
//ff:what 게이트 Context.Submission — 번역 글 단일 Target(Base nil 규약) + Prepare가 적재한 원문 Article, 언어쌍과 인스턴스 루트
package translation

import agate "github.com/park-jun-woo/abloq/pkg/gate"

// Submission is what one translation-quest submit carries through the gate:
// the assembled single-article target for the translation (Base nil —
// everything is judged as new), the parsed origin (default-language) article
// Prepare loaded for the parity comparison, the language pair and the
// instance root (the hugo-build rule builds the whole instance).
type Submission struct {
	Target     *agate.Target
	Origin     *agate.Article
	Article    string
	Root       string
	Lang       string
	OriginLang string
}
