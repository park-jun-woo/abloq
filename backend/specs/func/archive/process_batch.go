package archive

import (
	"encoding/json"

	parchive "github.com/park-jun-woo/abloq/pkg/archive"
)

// @func processBatch
// @error 500
// @description Execute up to limit pending receipts against Wayback SPN2, IndexNow and the GSC Indexing API with the site's submission credentials (IndexNow key, GSC SA JSON path — empty fields fall back to the instance-global env) and return the per-target results (done|failed|deferred) as receipt-upsert JSON — kind dispatch, per-target failure isolation and the GSC quota split live in pkg/archive

type ProcessBatchRequest struct {
	IndexNowKey   string
	GSCSAJSONPath string
	ReceiptsJSON  string
	Limit         int64
}

type ProcessBatchResponse struct {
	ResultsJSON []byte
	Processed   int64
	Done        int64
	Failed      int64
	Deferred    int64
}

// ProcessBatch is the thin @call wrapper around pkg/archive.ProcessBatchWith:
// JSON translation and status tallying only. The site row's credentials ride
// in as archive.Keys — a site's IndexNow submission must never go out under
// another site's key (multisite isolation), and an empty field keeps the
// single-site env fallback. The external HTTP work (and the env-driven base
// URLs) is entirely inside pkg/archive, which the abloq CLI shares
// (`abloq archive <url>`).
func ProcessBatch(req ProcessBatchRequest) (ProcessBatchResponse, error) {
	var pending []parchive.Pending
	if err := json.Unmarshal([]byte(req.ReceiptsJSON), &pending); err != nil {
		return ProcessBatchResponse{}, err
	}
	keys := parchive.Keys{IndexNowKey: req.IndexNowKey, GSCSAJSONPath: req.GSCSAJSONPath}
	results := parchive.ProcessBatchWith(keys, pending, req.Limit)
	// Marshal cannot fail: every result Request/Response is built by
	// pkg/archive (requestJSON/wrapResponse) and is valid JSON by construction.
	data, _ := json.Marshal(results)
	resp := ProcessBatchResponse{ResultsJSON: data, Processed: int64(len(results))}
	for _, r := range results {
		switch r.Status {
		case parchive.StatusDone:
			resp.Done++
		case parchive.StatusFailed:
			resp.Failed++
		case parchive.StatusDeferred:
			resp.Deferred++
		}
	}
	return resp, nil
}
