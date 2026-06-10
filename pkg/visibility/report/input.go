//ff:type feature=visibility type=schema topic=report
//ff:what 리포트 생성 입력 — ym, posts 인덱스, 현·전월 윈도 집계(크롤·GSC·인용), 큐 요약, 미지 봇, URL맵, 가중치
package report

import (
	"github.com/park-jun-woo/abloq/pkg/content"
	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
	"github.com/park-jun-woo/abloq/pkg/visibility/priority"
)

// Input carries everything Build needs. YM must already be resolved
// (ResolveYM); the Prev* slices hold the same aggregates over the previous
// month's window. URLs is the repository URL reverse map (cflog.BuildURLMap)
// for the GSC page attribution.
type Input struct {
	YM          string
	Posts       []content.Entry
	Bots        []BotSum
	PrevBots    []BotSum
	Pages       []PageSum
	PrevPages   []PageSum
	Cites       []CiteSum
	PrevCites   []CiteSum
	Queue       []QueueCount
	UnknownBots []UnknownBot
	URLs        map[string]cflog.Article
	Weights     priority.Weights
}
