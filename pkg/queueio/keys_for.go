//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 선언 언어 목록 × (section/slug) → 전 언어 조인 키 목록 — 스캐너 3종이 keys: 적재에 공유 (Phase018)
//ff:why 번역 재동기화가 전 언어의 lastmod 갱신을 일으키므로, 큐 파일 1개가 전 언어 키를 동반 발급해야 honest-lastmod 큐 등재 검사가 번역 커밋을 재차단하지 않는다 (Phase017 D4 해소)
package queueio

// KeysFor builds the gate-contract join key for every declared language.
// The order follows the blog.yaml languages declaration, so the serialized
// keys: block is deterministic for a given blog.
func KeysFor(langs []string, section, slug string) []string {
	if len(langs) == 0 {
		return nil
	}
	keys := make([]string, 0, len(langs))
	for _, lang := range langs {
		keys = append(keys, JoinKey(lang, section, slug))
	}
	return keys
}
