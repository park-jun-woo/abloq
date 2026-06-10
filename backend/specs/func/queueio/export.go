package queueio

import (
	"encoding/json"

	pqueueio "github.com/park-jun-woo/abloq/pkg/queueio"
)

// @func export
// @error 500
// @description Run one queue export cycle on the dedicated work clone (QUEUE_EXPORT_REPO_URL / QUEUE_EXPORT_WORKDIR): pull, detect agent-deleted files (consumed), write open items as quests/queue/*.yaml and push — git work, idempotent no-op commits and the deterministic serialization live in pkg/queueio, which Phase010/011 scanners reuse

type ExportRequest struct {
	OpenJSON     string
	ExportedJSON string
}

type ExportResponse struct {
	ExportedIdsJSON []byte
	ConsumedIdsJSON []byte
	Exported        int64
	Consumed        int64
}

// Export is the thin @call wrapper around pkg/queueio.Export: JSON
// translation only. Configuration (repo URL, work clone, commit author) comes
// from the environment so no path or credential flows through the API
// surface; the deploy key rides on GIT_SSH_COMMAND.
func Export(req ExportRequest) (ExportResponse, error) {
	cfg, err := pqueueio.ConfigFromEnv()
	if err != nil {
		return ExportResponse{}, err
	}
	open, err := pqueueio.DecodeRows([]byte(req.OpenJSON))
	if err != nil {
		return ExportResponse{}, err
	}
	exported, err := pqueueio.DecodeRows([]byte(req.ExportedJSON))
	if err != nil {
		return ExportResponse{}, err
	}
	res, err := pqueueio.Export(cfg, open, exported)
	if err != nil {
		return ExportResponse{}, err
	}
	// Marshal cannot fail: both inputs are []int64 (never nil — see rowIDs).
	exportedIDs, _ := json.Marshal(res.ExportedIDs)
	consumedIDs, _ := json.Marshal(res.ConsumedIDs)
	return ExportResponse{
		ExportedIdsJSON: exportedIDs,
		ConsumedIdsJSON: consumedIDs,
		Exported:        int64(len(res.ExportedIDs)),
		Consumed:        int64(len(res.ConsumedIDs)),
	}, nil
}
