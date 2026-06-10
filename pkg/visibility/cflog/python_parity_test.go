//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what analyze-stats.py 원시 분류 대조 — 픽스처 전체의 무필터 봇별 원시 카운터를 python 1회 실행 골든값과 비교 (BOTPAT 통과 사전 교집합 봇 한정)
//ff:why 이식 정확성은 원시 분류 단계로 정의한다: crawl_hits(2xx/304·매핑 성공분)와 직접 비교하면 기준이 달라 무의미. GoogleOther는 BOTPAT 미통과 사문 항목이라 교집합에서 제외 — python은 사람으로 센다 (Phase012)
package cflog

import (
	"testing"
)

// TestPythonParity compares the unfiltered per-bot raw counters over every
// committed testdata/logs fixture against the analyze-stats.py bots
// counters, restricted to the dictionary∩BOT_NAMES bots that actually pass
// the python is_bot BOTPAT gate (GoogleOther does not — dead entry).
//
// Golden values were produced by one python run (the Go test is hermetic):
//
//	W=$(mktemp -d); cp ~/.clari/repos/parkjunwoo/deploy/scripts/analyze-stats.py "$W/"
//	mkdir "$W/deploy"; ln -s <abloq>/pkg/visibility/cflog/testdata/logs "$W/deploy/logs"
//	(cd "$W" && python3 analyze-stats.py 2026-06-01 2026-06-02 2026-06-03)
//	# TARGET = every KST date the fixtures touch, last UTC day +1 included
//	# 봇상세 summed over the days:
//	# {'GPTBot': 4, 'ClaudeBot': 4, 'ChatGPT-User': 2, 'PerplexityBot': 2,
//	#  'PetalBot': 1, '(빈 UA)': 1, 'Amazonbot': 1, 'Bytespider': 1,
//	#  'OAI-SearchBot': 1, '기타봇': 1}
func TestPythonParity(t *testing.T) {
	src := DirSource{Root: "testdata/logs"}
	keys, err := src.List("", "")
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	agg, err := IngestKeys(src, nil, keys)
	if err != nil {
		t.Fatalf("IngestKeys: %v", err)
	}
	golden := map[string]int64{
		"GPTBot":        4,
		"ClaudeBot":     4,
		"ChatGPT-User":  2,
		"PerplexityBot": 2,
		"Amazonbot":     1,
		"Bytespider":    1,
		"OAI-SearchBot": 1,
	}
	for bot, want := range golden {
		if got := agg.Raw[bot]; got != want {
			t.Errorf("raw counter %s = %d, python golden %d", bot, got, want)
		}
	}
}
