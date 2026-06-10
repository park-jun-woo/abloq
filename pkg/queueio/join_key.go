//ff:func feature=queueio type=generator control=sequence
//ff:what 게이트 계약 조인 키 조립 — <lang>/<section>/<slug> 원문 그대로 (honest-lastmod queueAllows가 부분문자열 매칭)
package queueio

// JoinKey builds the gate-contract article key. pkg/gate's queueAllows
// matches this exact substring inside a queue file, so the serialization
// must embed it verbatim.
func JoinKey(lang, section, slug string) string {
	return lang + "/" + section + "/" + slug
}
