//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what 선별된 로그 키들을 순서대로 받아 파싱·누적 — 수집 코어, 커서를 모름 (CLI 단발 분석과 백엔드 수집이 공유)
package cflog

// IngestKeys streams every key from the source through the parser into one
// accumulator. It knows nothing about cursors or margins — the backend's
// Collect filters keys first, the CLI feeds every log key.
func IngestKeys(src Source, urls map[string]Article, keys []string) (*Agg, error) {
	agg := NewAgg(urls)
	for _, key := range keys {
		rc, err := src.Get(key)
		if err != nil {
			return nil, err
		}
		recs, err := parseRecords(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}
		for _, rec := range recs {
			agg.Add(rec)
		}
	}
	return agg, nil
}
