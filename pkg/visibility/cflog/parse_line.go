//ff:func feature=visibility type=parser control=sequence topic=crawl
//ff:what CF 로그 1행 파싱 — # 헤더/빈 행/11필드 미만/시각 파싱 실패는 버리고, URI·UA를 unquote한 Record 반환
//ff:why analyze-stats.py L49-61과 같은 탈락 기준(주석, len<11, ValueError continue)을 그대로 승계해야 원시 카운터 대조가 일치한다 (Phase012)
package cflog

import (
	"strings"
	"time"
)

// parseLine parses one CloudFront standard-log line (W3C extended, tab
// separated). The second return is false for header lines, blank lines,
// short rows and unparseable timestamps — the same drops analyze-stats.py
// applies.
func parseLine(line string) (Record, bool) {
	if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
		return Record{}, false
	}
	p := strings.Split(strings.TrimRight(line, "\n"), "\t")
	if len(p) < 11 {
		return Record{}, false
	}
	when, err := time.Parse("2006-01-02 15:04:05", p[0]+" "+p[1])
	if err != nil {
		return Record{}, false
	}
	return Record{
		When:   when.UTC(),
		URI:    unquote(p[7]),
		Status: p[8],
		UA:     unquote(p[10]),
	}, true
}
