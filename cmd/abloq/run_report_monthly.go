//ff:func feature=cli type=command control=sequence topic=report
//ff:what monthly 리포트 실행 본체 — 저장소 인덱스 + (선택) CF 로그의 윈도 집계로 부분 리포트 markdown 출력, 인용·GSC 계층은 원리적 불가를 명시
//ff:why CLI는 무상태다: 인용 계층은 저장된 샘플 시계열 없이는 원리적으로 불가(샘플은 실행 시점 기록)고 GSC 직조회는 범위 외 보수 선택 — 부분 리포트임을 출력에 정직하게 박는다 (Phase014)
package main

import (
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
	"github.com/park-jun-woo/abloq/pkg/visibility/report"
)

// runReportMonthly renders the stateless partial report: the posts index
// plus the crawl layer aggregated from the log source for the ym window
// (and the previous window for the trend). Citation and GSC columns stay
// empty and the header says the report is partial.
func runReportMonthly(out io.Writer, dir, ym, source string) error {
	b, err := loadValidBlog(out, dir)
	if err != nil {
		return err
	}
	resolved, err := report.ResolveYM(ym, time.Now().UTC())
	if err != nil {
		return err
	}
	posts := content.IndexEntries(dir, b)
	bots, prevBots, err := collectLogBotSums(dir, source, resolved, b)
	if err != nil {
		return err
	}
	r := report.Build(report.Input{
		YM:    resolved,
		Posts: posts,
		Bots:  bots, PrevBots: prevBots,
		Weights: priority.WeightsOf(b.Geo.PriorityWeights),
		URLs:    map[string]cflog.Article{},
	})
	fmt.Fprintln(out, "<!-- PARTIAL REPORT (abloq CLI, stateless): crawl layer only. -->")
	fmt.Fprintln(out, "<!-- Citation needs the stored sample time series and GSC the snapshot store — both live in abloqd (POST /reports/monthly). -->")
	fmt.Fprint(out, report.Markdown(r))
	return nil
}
