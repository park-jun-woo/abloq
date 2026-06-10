//ff:type feature=archive type=schema
//ff:what 대기 영수증 1건 — ReceiptAggPendingJson 산출(deploy_id/kind/target + posts join의 date/lastmod), GSC 우선순위 입력
package archive

// Pending is one pending receipt as aggregated by ReceiptAggPendingJson.
// Date/Lastmod come from the posts index join and drive the GSC priority
// (new post = date == lastmod); both are empty when the target URL is not
// in the index.
type Pending struct {
	DeployID string `json:"deploy_id"`
	Kind     string `json:"kind"`
	Target   string `json:"target"`
	Date     string `json:"date"`
	Lastmod  string `json:"lastmod"`
}
