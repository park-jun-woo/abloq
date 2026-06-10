//ff:func feature=archive type=client control=sequence
//ff:what wrapResponse가 JSON 본문은 원형 임베드, 텍스트는 문자열, 2KB 초과는 절단하는지 검증
package archive

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestWrapResponse(t *testing.T) {
	var got map[string]any
	if err := json.Unmarshal(wrapResponse(200, []byte(`{"job_id":"j1"}`)), &got); err != nil {
		t.Fatalf("unmarshal json-body wrap: %v", err)
	}
	if got["status_code"].(float64) != 200 {
		t.Errorf("status_code = %v, want 200", got["status_code"])
	}
	if body, ok := got["body"].(map[string]any); !ok || body["job_id"] != "j1" {
		t.Errorf("body = %v, want embedded {job_id: j1}", got["body"])
	}

	if err := json.Unmarshal(wrapResponse(0, []byte("dial refused")), &got); err != nil {
		t.Fatalf("unmarshal text-body wrap: %v", err)
	}
	if got["body"] != "dial refused" {
		t.Errorf("text body = %v, want string \"dial refused\"", got["body"])
	}

	long := strings.Repeat("x", 5000)
	if err := json.Unmarshal(wrapResponse(500, []byte(long)), &got); err != nil {
		t.Fatalf("unmarshal truncated wrap: %v", err)
	}
	if len(got["body"].(string)) != 2048 {
		t.Errorf("truncated body length = %d, want 2048", len(got["body"].(string)))
	}
}
