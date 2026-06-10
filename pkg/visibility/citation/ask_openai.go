//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what OpenAI 질의 1발 — Responses API에 web_search 도구로 POST, 메시지 출력의 url_citation URL을 수집
package citation

import (
	"encoding/json"
	"net/http"
)

// askOpenAI asks one query through the OpenAI Responses API with the
// web_search tool and returns the url_citation annotation URLs. base/model
// are env-bound by EnginesFromEnv (OPENAI_BASE_URL / OPENAI_MODEL).
func askOpenAI(base, key, model, query string) ([]string, error) {
	header := http.Header{"Authorization": {"Bearer " + key}}
	body, err := postJSON(base+"/v1/responses", header, map[string]any{
		"model": model,
		"tools": []map[string]string{{"type": "web_search"}},
		"input": query,
	})
	if err != nil {
		return nil, err
	}
	var parsed oaiResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	var urls []string
	for _, out := range parsed.Output {
		urls = append(urls, urlCitations(out.Content)...)
	}
	return urls, nil
}
