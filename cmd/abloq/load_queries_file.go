//ff:func feature=cli type=parser control=iteration dimension=1 topic=citation
//ff:what 질의 파일(yaml|json) 파싱 — {id, query_text} 목록으로 디코드, 빈 목록·빈 질의문은 에러
//ff:why CLI는 DB 미접속: citation_queries는 백엔드 소관이고 CLI 입력은 인자 파일이다 — 같은 JSON 키라 백엔드 jsonb_agg 출력도 그대로 먹는다 (Phase013)
package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/abloq/pkg/visibility/citation"
)

// loadQueriesFile reads the --queries file: a YAML (or JSON — YAML 1.2
// superset) list of {id, query_text} entries, the same keys the backend's
// citation_queries aggregate emits.
func loadQueriesFile(path string) ([]citation.Query, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var queries []citation.Query
	if err := yaml.Unmarshal(data, &queries); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if len(queries) == 0 {
		return nil, fmt.Errorf("%s: no queries", path)
	}
	for i, q := range queries {
		if q.Text == "" {
			return nil, fmt.Errorf("%s: entry %d has no query_text", path, i)
		}
	}
	return queries, nil
}
