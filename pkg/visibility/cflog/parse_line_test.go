//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what parseLine이 정상 행을 디코드하고 # 헤더·빈 행·11필드 미만·시각 불량을 버리는지 검증 (analyze-stats 탈락 기준 승계)
package cflog

import (
	"testing"
	"time"
)

func TestParseLine(t *testing.T) {
	good := "2026-06-01\t12:00:01\tICN54\t1024\t203.0.113.10\tGET\tblog.test\t/tech/post%2Da/\t200\t-\tBot%20UA/1.0\t-\t-\tHit\trid"
	rec, ok := parseLine(good)
	if !ok {
		t.Fatalf("parseLine dropped a good line")
	}
	if rec.URI != "/tech/post-a/" || rec.UA != "Bot UA/1.0" || rec.Status != "200" {
		t.Errorf("rec = %+v", rec)
	}
	if !rec.When.Equal(time.Date(2026, 6, 1, 12, 0, 1, 0, time.UTC)) {
		t.Errorf("When = %v", rec.When)
	}
	drops := []string{
		"#Version: 1.0",
		"",
		"   ",
		"2026-06-01\t12:00:01\tshort",
		"not-a-date\t12:00:01\tICN54\t1024\tip\tGET\thost\t/\t200\t-\tua\t-\t-\tHit\trid",
	}
	for _, line := range drops {
		if _, ok := parseLine(line); ok {
			t.Errorf("parseLine(%q) accepted, want drop", line)
		}
	}
}
