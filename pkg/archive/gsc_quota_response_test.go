//ff:func feature=archive type=client control=sequence
//ff:what gscQuotaResponse가 wrapResponse에 quota_exceeded: true 마커를 더하는지 검증
package archive

import (
	"encoding/json"
	"testing"
)

func TestGscQuotaResponse(t *testing.T) {
	var got map[string]any
	if err := json.Unmarshal(gscQuotaResponse(429, []byte(`{"error":"quota"}`)), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got["quota_exceeded"] != true {
		t.Errorf("quota_exceeded = %v, want true", got["quota_exceeded"])
	}
	if got["status_code"].(float64) != 429 {
		t.Errorf("status_code = %v, want 429", got["status_code"])
	}
	if _, ok := got["body"]; !ok {
		t.Error("body evidence missing")
	}
}
