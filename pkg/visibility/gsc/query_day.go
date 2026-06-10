//ff:func feature=visibility type=client control=iteration dimension=1 topic=gsc
//ff:what Search Analytics 1일 조회 — dimensions=[page]로 페이지별 노출·클릭·평균순위를 Snapshot 행으로 환원
package gsc

import (
	"encoding/json"
	"math"
	"net/url"
)

// QueryDay fetches one closed day from the Search Analytics API: every page
// of the property with its impressions, clicks and average position. The
// base is env-overridable upstream (GSC_SEARCH_API_BASE) so the Hurl stub
// can intercept it.
func QueryDay(base, token, site, date string) ([]Snapshot, error) {
	endpoint := base + "/webmasters/v3/sites/" + url.PathEscape(site) + "/searchAnalytics/query"
	body, err := postJSON(endpoint, token, map[string]any{
		"startDate":  date,
		"endDate":    date,
		"dimensions": []string{"page"},
		"rowLimit":   25000,
	})
	if err != nil {
		return nil, err
	}
	var parsed struct {
		Rows []struct {
			Keys        []string `json:"keys"`
			Clicks      float64  `json:"clicks"`
			Impressions float64  `json:"impressions"`
			Position    float64  `json:"position"`
		} `json:"rows"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	var rows []Snapshot
	for _, r := range parsed.Rows {
		if len(r.Keys) == 0 {
			continue
		}
		rows = append(rows, Snapshot{
			SnapDate:    date,
			Page:        r.Keys[0],
			Impressions: int64(math.Round(r.Impressions)),
			Clicks:      int64(math.Round(r.Clicks)),
			AvgPosition: r.Position,
		})
	}
	return rows, nil
}
