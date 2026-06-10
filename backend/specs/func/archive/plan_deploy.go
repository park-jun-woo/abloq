package archive

import (
	"encoding/json"

	parchive "github.com/park-jun-woo/abloq/pkg/archive"
)

// @func planDeploy
// @error 500
// @description Plan the pending archive receipts of one deploy — every (kind, target) pair of the changed URLs minus the pairs that already hold a receipt for this deploy_id (idempotent re-webhook)

type PlanDeployRequest struct {
	DeployID     string
	Changed      []string
	ReceiptsJSON string
}

type PlanDeployResponse struct {
	ItemsJSON []byte
	Planned   int64
}

// PlanDeploy is the thin @call wrapper around pkg/archive.PlanDeploy: it only
// translates between the jsonb_agg/jsonb_array_elements JSON contracts of the
// receipts queries and the typed pkg/archive API. No external call happens
// here — the webhook stays fast and side-effect free.
func PlanDeploy(req PlanDeployRequest) (PlanDeployResponse, error) {
	var existing []parchive.Existing
	if err := json.Unmarshal([]byte(req.ReceiptsJSON), &existing); err != nil {
		return PlanDeployResponse{}, err
	}
	items := parchive.PlanDeploy(req.DeployID, req.Changed, existing)
	// Marshal cannot fail: Item carries only strings and the literal `{}`
	// raw JSON at the pending stage.
	data, _ := json.Marshal(items)
	return PlanDeployResponse{ItemsJSON: data, Planned: int64(len(items))}, nil
}
