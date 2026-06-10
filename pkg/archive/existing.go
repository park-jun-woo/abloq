//ff:type feature=archive type=schema
//ff:what 기존 영수증 (kind, target) 페어 1건 — ReceiptAggByDeployJson 산출, 웹훅 멱등 필터 입력
package archive

// Existing is one already-receipted (kind, target) pair of a deploy, as
// aggregated by ReceiptAggByDeployJson.
type Existing struct {
	Kind   string `json:"kind"`
	Target string `json:"target"`
}
