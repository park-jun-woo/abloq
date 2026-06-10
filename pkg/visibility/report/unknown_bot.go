//ff:type feature=visibility type=schema topic=report
//ff:what 미지 봇 후보 1행 — UA와 누적 히트, UnknownBot.AggUasJson 행과 1:1 (사전 증보 입력)
package report

// UnknownBot is one unknown-bot candidate row: a UA the crawl ingest saw
// that is not in the pkg/bots dictionary. The report surfaces them as the
// operator's dictionary-update input.
type UnknownBot struct {
	UA   string `json:"ua"`
	Hits int64  `json:"hits"`
}
