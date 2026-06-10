//ff:type feature=archive type=schema
//ff:what 영수증 업서트 항목 1건 — deploy_id/kind/target/request/response/status, receipts 테이블 1행과 1:1
//ff:why JSON 태그가 ReceiptUpsertFromJson(jsonb_array_elements)의 키와 일치해야 한다 — 백엔드 @call 래퍼가 이 배열을 그대로 배치 업서트에 먹인다 (Phase008)
package archive

import "encoding/json"

// Receipt kinds — the fixed submission fan-out of one changed URL.
const (
	KindWayback  = "wayback"
	KindIndexNow = "indexnow"
	KindGSCIndex = "gsc_index"
)

// Receipt statuses — pending → done | failed | deferred.
const (
	StatusPending  = "pending"
	StatusDone     = "done"
	StatusFailed   = "failed"
	StatusDeferred = "deferred"
)

// Kinds lists every receipt kind in submission order.
var Kinds = []string{KindWayback, KindIndexNow, KindGSCIndex}

// Item is one receipt row for the batch upsert (ReceiptUpsertFromJson).
type Item struct {
	DeployID string          `json:"deploy_id"`
	Kind     string          `json:"kind"`
	Target   string          `json:"target"`
	Request  json.RawMessage `json:"request"`
	Response json.RawMessage `json:"response"`
	Status   string          `json:"status"`
}
