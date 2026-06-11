//ff:func feature=queueio type=generator control=sequence
//ff:what 게이트 계약 조인 키 조립 — <lang>/<section>/<slug> 원문 그대로 (honest-lastmod queueAllows가 인용형 정확 매칭)
package queueio

// JoinKey builds the gate-contract article key. pkg/gate's queueAllows
// matches this key exactly in its quoted form (the strconv.Quote framing the
// serialization puts on every key:/keys: line), so the serialization must
// embed it verbatim and quoted.
func JoinKey(lang, section, slug string) string {
	return lang + "/" + section + "/" + slug
}
