//ff:func feature=visibility type=client control=sequence topic=gsc
//ff:what URL Inspection 1건 — inspectionUrl·siteUrl POST, indexStatusResult의 verdict·coverageState 요약으로 환원
package gsc

import "encoding/json"

// inspectOne asks the URL Inspection API about one URL and reduces the
// answer to its index-status verdict summary.
func inspectOne(endpoint, token, site, u string) (Inspection, error) {
	body, err := postJSON(endpoint, token, map[string]string{
		"inspectionUrl": u,
		"siteUrl":       site,
	})
	if err != nil {
		return Inspection{}, err
	}
	var parsed struct {
		InspectionResult struct {
			IndexStatusResult struct {
				Verdict       string `json:"verdict"`
				CoverageState string `json:"coverageState"`
			} `json:"indexStatusResult"`
		} `json:"inspectionResult"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return Inspection{}, err
	}
	return Inspection{
		URL:           u,
		Verdict:       parsed.InspectionResult.IndexStatusResult.Verdict,
		CoverageState: parsed.InspectionResult.IndexStatusResult.CoverageState,
	}, nil
}
