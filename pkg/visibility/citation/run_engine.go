//ff:func feature=visibility type=client control=iteration dimension=1 topic=citation
//ff:what 엔진 1종 1회전 — active 질의를 id 순으로 budget만큼 실행, 호출 사이 interval 대기(엔진별 rate limit), 실패도 에러 근거 샘플로 기록
//ff:why 샘플은 추세 기록 자체가 목적 — 엔진 에러를 삼키면 시계열에 가짜 공백이 생긴다. 에러도 cited=false + {"error":...}로 남긴다 (Phase013)
package citation

import "time"

// runEngine samples one engine over the first budget queries: ask, match the
// own domain, record. interval throttles consecutive calls (the per-engine
// rate limit); a failed ask still yields a sample with the error as
// evidence.
func runEngine(e Engine, queries []Query, budget int, domain string, interval time.Duration) []Sample {
	if budget > len(queries) {
		budget = len(queries)
	}
	var samples []Sample
	for i := 0; i < budget; i++ {
		if i > 0 {
			time.Sleep(interval)
		}
		urls, err := e.Ask(queries[i].Text)
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		matched := Matched(domain, urls)
		samples = append(samples, Sample{
			CitationQueriesID: queries[i].ID,
			Engine:            e.Name,
			Cited:             errMsg == "" && len(matched) > 0,
			Evidence:          evidenceJSON(matched, errMsg),
			ExtractorVersion:  ExtractorVersion,
		})
	}
	return samples
}
