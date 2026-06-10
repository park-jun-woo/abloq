//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what Anthropic 질의 1발 — Messages API에 web_search 도구로 POST, content[].citations의 URL을 수집
package citation

import (
	"encoding/json"
	"net/http"
)

// askAnthropic asks one query through the Anthropic Messages API with the
// web_search server tool and returns the citation URLs of the text blocks.
// base/model are env-bound by EnginesFromEnv (ANTHROPIC_BASE_URL /
// ANTHROPIC_MODEL).
func askAnthropic(base, key, model, query string) ([]string, error) {
	header := http.Header{
		"X-Api-Key":         {key},
		"Anthropic-Version": {"2023-06-01"},
	}
	body, err := postJSON(base+"/v1/messages", header, map[string]any{
		"model":      model,
		"max_tokens": 1024,
		"messages":   []map[string]string{{"role": "user", "content": query}},
		"tools":      []map[string]any{{"type": "web_search_20250305", "name": "web_search", "max_uses": 3}},
	})
	if err != nil {
		return nil, err
	}
	var parsed anthropicResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	// Empty URLs (non-web citation kinds) ride along harmlessly — the
	// domain matcher never matches them.
	var urls []string
	for _, block := range parsed.Content {
		for _, c := range block.Citations {
			urls = append(urls, c.URL)
		}
	}
	return urls, nil
}
