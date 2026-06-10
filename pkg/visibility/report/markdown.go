//ff:func feature=visibility type=generator control=iteration dimension=1 topic=report
//ff:what 리포트 → 공개용 markdown — 글별 표·전월 대비(n/a)·미지 봇·큐 요약, 민감 정보 없이 그대로 발행 가능 (§7 공개 증거 시퀀스)
package report

import (
	"fmt"
	"strings"
)

// Markdown renders the human-readable, publication-ready report. It carries
// only ym-anchored aggregates — no secret, no credential, no clock value —
// so the file can be committed to the public blog repository verbatim (§7
// public evidence sequence). The previous-month column reads "n/a" on the
// first month.
func Markdown(r Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Visibility report %s\n\n", r.YM)
	fmt.Fprintf(&b, "Window: %s .. %s (30 days ending on the last day of %s, both ends included).\n\n", r.WindowFrom, r.WindowTo, r.YM)
	b.WriteString("## Articles\n\n")
	b.WriteString("| article | date | training | search | fetch | md | gsc impressions | gsc clicks | cited | priority |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, row := range r.Rows {
		fmt.Fprintf(&b, "| %s/%s/%s | %s | %d | %d | %d | %d | %d | %d | %d | %d |\n",
			row.Lang, row.Section, row.Slug, row.Date,
			row.Training, row.Search, row.Fetch, row.MDHits,
			row.Impressions, row.Clicks, row.Cited, row.Priority)
	}
	fmt.Fprintf(&b, "\n## Month over month (%s vs %s)\n\n", r.YM, r.PrevYM)
	fmt.Fprintf(&b, "| metric | %s | %s |\n|---|---:|---:|\n", r.YM, r.PrevYM)
	fmt.Fprintf(&b, "| crawl hits | %d | %s |\n", r.Totals.CrawlHits, prevCell(r, r.PrevTotals.CrawlHits))
	fmt.Fprintf(&b, "| md hits | %d | %s |\n", r.Totals.MDHits, prevCell(r, r.PrevTotals.MDHits))
	fmt.Fprintf(&b, "| gsc impressions | %d | %s |\n", r.Totals.Impressions, prevCell(r, r.PrevTotals.Impressions))
	fmt.Fprintf(&b, "| gsc clicks | %d | %s |\n", r.Totals.Clicks, prevCell(r, r.PrevTotals.Clicks))
	fmt.Fprintf(&b, "| cited samples | %d | %s |\n", r.Totals.Cited, prevCell(r, r.PrevTotals.Cited))
	b.WriteString("\n## Unknown bots\n\n")
	b.WriteString(unknownSection(r.UnknownBots))
	b.WriteString("\n## Queue intake\n\n")
	b.WriteString(queueSection(r.Queue))
	return b.String()
}
