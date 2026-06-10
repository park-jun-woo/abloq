//ff:func feature=visibility type=client control=sequence topic=citation
//ff:what Perplexity 질의 1발 — chat/completions POST, 응답 최상위 citations 배열(구조화 필드)을 그대로 반환
package citation

import (
	"encoding/json"
	"net/http"
)

// askPerplexity asks one query through the Perplexity chat completions API
// and returns the structured top-level citations array. base/model are
// env-bound by EnginesFromEnv (PERPLEXITY_BASE_URL / PERPLEXITY_MODEL).
func askPerplexity(base, key, model, query string) ([]string, error) {
	header := http.Header{"Authorization": {"Bearer " + key}}
	body, err := postJSON(base+"/chat/completions", header, map[string]any{
		"model":    model,
		"messages": []map[string]string{{"role": "user", "content": query}},
	})
	if err != nil {
		return nil, err
	}
	var parsed struct {
		Citations []string `json:"citations"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	return parsed.Citations, nil
}
