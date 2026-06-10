//ff:type feature=visibility type=schema topic=report
//ff:what 월간 가시성 리포트 — ym·글별 표·윈도 합계·전월 합계(데이터 유무 플래그)·큐 요약·미지 봇, JSON 산출의 원형
package report

// Report is the monthly visibility report: the per-article table, the
// window totals against the previous window (PrevHasData false = first
// month, rendered "n/a"), the queue intake summary and the unknown-bot
// candidates. Every value derives from ym-anchored DB aggregates — no
// clock-derived value appears, so regeneration is byte-identical.
type Report struct {
	YM          string       `json:"ym"`
	PrevYM      string       `json:"prev_ym"`
	WindowFrom  string       `json:"window_from"`
	WindowTo    string       `json:"window_to"`
	Rows        []Row        `json:"articles"`
	Totals      Totals       `json:"totals"`
	PrevTotals  Totals       `json:"prev_totals"`
	PrevHasData bool         `json:"prev_has_data"`
	Queue       []QueueCount `json:"queue"`
	UnknownBots []UnknownBot `json:"unknown_bots"`
}
