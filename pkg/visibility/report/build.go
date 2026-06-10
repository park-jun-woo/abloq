//ff:func feature=visibility type=generator control=sequence topic=report
//ff:what 리포트 조립 — 분류·귀속·결합·점수·합계·전월 비교를 한 번에, 같은 Input이면 바이트 동일 산출의 원형
package report

// Build assembles the monthly report from resolved inputs. Deterministic:
// identical Input always yields an identical Report (and so identical
// markdown/JSON), which the idempotent git publication rests on.
func Build(in Input) Report {
	from, to, _ := WindowDates(in.YM)
	bots := BotTotals(in.Bots)
	pages := PageTotals(in.Pages, in.URLs)
	cites := CiteHits(in.Cites)
	queue := in.Queue
	if queue == nil {
		queue = []QueueCount{}
	}
	unknown := in.UnknownBots
	if unknown == nil {
		unknown = []UnknownBot{}
	}
	return Report{
		YM:          in.YM,
		PrevYM:      PrevYM(in.YM),
		WindowFrom:  from,
		WindowTo:    to,
		Rows:        rows(in.Posts, bots, pages, cites, in.Weights),
		Totals:      TotalsOf(bots, pages, cites),
		PrevTotals:  TotalsOf(BotTotals(in.PrevBots), PageTotals(in.PrevPages, in.URLs), CiteHits(in.PrevCites)),
		PrevHasData: len(in.PrevBots) > 0 || len(in.PrevPages) > 0 || len(in.PrevCites) > 0,
		Queue:       queue,
		UnknownBots: unknown,
	}
}
