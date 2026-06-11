//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what Seed 시점 고정된 payload candidates JSON → 허용 추가 경로(기본 언어 후보 글의 플랫/번들 2형) — queue-scope 확장분
//ff:why no-isolated-post 해소는 후보 글에서 들어오는 링크가 필수다 — 백엔드 발급 payload 유래의 확장이라 치즈가 아니며, 이게 없으면 고립 위반 아이템은 구조적 해소 불가 → MaxTries 교착 (Phase018 계획)
package cluster

import (
	"encoding/json"
	"fmt"

	scancluster "github.com/park-jun-woo/abloq/pkg/scan/cluster"
)

// candidatePaths decodes the frozen queue payload's candidates entry and
// expands each suggestion into its article path forms (flat and bundle) in
// the item's language — the cluster graph is default-language only, so the
// item language is the candidates' language. An absent entry yields nil.
func candidatePaths(payload map[string]string, lang string) ([]string, error) {
	raw, ok := payload["candidates"]
	if !ok {
		return nil, nil
	}
	var cands []scancluster.Candidate
	if err := json.Unmarshal([]byte(raw), &cands); err != nil {
		return nil, fmt.Errorf("queue payload candidates: %w", err)
	}
	var paths []string
	for _, c := range cands {
		paths = append(paths, "content/"+lang+"/"+c.Section+"/"+c.Slug+".md")
		paths = append(paths, "content/"+lang+"/"+c.Section+"/"+c.Slug+"/index.md")
	}
	return paths, nil
}
