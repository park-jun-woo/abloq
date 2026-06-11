//ff:func feature=quest type=generator control=iteration dimension=1 topic=queue
//ff:what 전 언어 키 목록 → queue-scope 허용 경로 집합 — 키마다 플랫/번들 글 경로 + insight 사이드카 2형, 큐 파일 자신은 불포함
//ff:why 허용 집합은 대상 글과 그 번역본(전 언어)·insight 사이드카까지다 — 큐 파일은 게이트 시점 무변경이어야 하므로(삭제는 ② 커밋에서만) 집합에서 제외한다 (Phase018 계획)
package common

// AllowedPaths expands the per-language join keys into the queue-scope
// allowed file set: for each key both article forms (flat and bundle) plus
// both insight sidecar forms (Phase015 convention). The queue file itself is
// deliberately absent — it must stay untouched until the consumption-signal
// commit after the gate.
func AllowedPaths(keys []string) map[string]bool {
	allowed := make(map[string]bool, len(keys)*4)
	for _, k := range keys {
		allowed["content/"+k+".md"] = true
		allowed["content/"+k+"/index.md"] = true
		allowed["content/"+k+".insight.yaml"] = true
		allowed["content/"+k+"/insight.yaml"] = true
	}
	return allowed
}
