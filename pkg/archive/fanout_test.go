//ff:func feature=archive type=client control=iteration dimension=1
//ff:what fanoutItems가 공유 response/status를 항목별 영수증으로 전개하고 per-target request를 만드는지 검증
package archive

import (
	"encoding/json"
	"testing"
)

func TestFanoutItems(t *testing.T) {
	pending := []Pending{
		{DeployID: "dep-1", Kind: KindIndexNow, Target: "https://a/"},
		{DeployID: "dep-1", Kind: KindIndexNow, Target: "https://b/"},
	}
	resp := json.RawMessage(`{"status_code":200}`)
	items := fanoutItems(pending, "https://api.indexnow.org/indexnow", resp, StatusDone)
	if len(items) != 2 {
		t.Fatalf("len = %d, want 2", len(items))
	}
	for i, item := range items {
		if item.Status != StatusDone || string(item.Response) != string(resp) {
			t.Errorf("item %d: status/response not shared: %+v", i, item)
		}
		var req map[string]string
		if err := json.Unmarshal(item.Request, &req); err != nil || req["url"] != pending[i].Target {
			t.Errorf("item %d: request url = %v (err=%v), want per-target", i, req, err)
		}
	}
	if fanoutItems(nil, "e", resp, StatusDone) == nil {
		t.Log("nil input yields empty slice — acceptable")
	}
}
