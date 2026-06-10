//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what IngestKeys가 받은 키만 순서대로 파싱·누적하고 없는 키는 에러인지 검증 — 커서를 모르는 수집 코어
package cflog

import "testing"

func TestIngestKeys(t *testing.T) {
	src := DirSource{Root: "testdata/logs"}
	agg, err := IngestKeys(src, nil, []string{"E123ABC.2026-06-01-22.bbbb2222.gz"})
	if err != nil {
		t.Fatalf("IngestKeys: %v", err)
	}
	if agg.Raw["Amazonbot"] != 1 || agg.Raw["Bytespider"] != 1 || agg.Raw["GPTBot"] != 0 {
		t.Errorf("Raw = %v, want only the hour-22 bots", agg.Raw)
	}
	if _, err := IngestKeys(src, nil, []string{"missing.gz"}); err == nil {
		t.Errorf("missing key accepted")
	}
}
