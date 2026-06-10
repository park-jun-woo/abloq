//ff:type feature=visibility type=schema topic=crawl
//ff:what 미지 봇 후보 1행 — UA 유니크 키, 히트 수, 최초/최종 목격 시각(UTC RFC3339), unknown_bots 1행 대응
//ff:why 사전 갱신 입력용 — JSON 키가 unknown_bots 컬럼명과 일치해야 백엔드 jsonb 업서트가 그대로 통한다 (Phase012)
package cflog

// UnknownRow is one unknown-bot candidate: a UA that is not in the bot
// dictionary but matches the bot heuristic. FirstSeen/LastSeen are UTC
// RFC3339 timestamps of the earliest/latest sighting in this ingest.
type UnknownRow struct {
	UA        string `json:"ua"`
	Hits      int64  `json:"hits"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}
