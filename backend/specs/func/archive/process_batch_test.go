package archive

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessBatch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{"ok":true}`)); err != nil {
			t.Errorf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("WAYBACK_BASE_URL", srv.URL)
	t.Setenv("INDEXNOW_ENDPOINT", srv.URL+"/indexnow")
	t.Setenv("INDEXNOW_KEY", "")
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")
	t.Setenv("GSC_DAILY_QUOTA", "1")

	receipts := `[
	  {"deploy_id":"d","kind":"wayback","target":"https://blog.example.com/p/","date":"","lastmod":""},
	  {"deploy_id":"d","kind":"indexnow","target":"https://blog.example.com/p/","date":"","lastmod":""},
	  {"deploy_id":"d","kind":"gsc_index","target":"https://blog.example.com/p/","date":"","lastmod":""},
	  {"deploy_id":"d","kind":"gsc_index","target":"https://blog.example.com/q/","date":"","lastmod":""}
	]`
	resp, err := ProcessBatch(ProcessBatchRequest{IndexNowKey: "k123", ReceiptsJSON: receipts, Limit: 100})
	if err != nil {
		t.Fatalf("ProcessBatch: %v", err)
	}
	// wayback + indexnow succeed against the stub; gsc has no SA creds, so the
	// quota head (1) fails and the quota tail (1) defers.
	if resp.Processed != 4 || resp.Done != 2 || resp.Failed != 1 || resp.Deferred != 1 {
		t.Errorf("tally = %d/%d/%d/%d (processed/done/failed/deferred), want 4/2/1/1",
			resp.Processed, resp.Done, resp.Failed, resp.Deferred)
	}

	empty, err := ProcessBatch(ProcessBatchRequest{ReceiptsJSON: "[]", Limit: 10})
	if err != nil || empty.Processed != 0 {
		t.Errorf("empty batch = %+v, %v, want no-op", empty, err)
	}

	if _, err := ProcessBatch(ProcessBatchRequest{ReceiptsJSON: "not json", Limit: 1}); err == nil {
		t.Error("invalid receipts JSON must fail")
	}
}
