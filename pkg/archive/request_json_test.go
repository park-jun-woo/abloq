//ff:func feature=archive type=client control=sequence
//ff:what requestJSON이 endpoint와 url만 기록하는지(자격증명 무기록) 검증
package archive

import (
	"encoding/json"
	"testing"
)

func TestRequestJSON(t *testing.T) {
	var got map[string]string
	if err := json.Unmarshal(requestJSON("https://api.example.com/save", "https://blog.example.com/p/"), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got["endpoint"] != "https://api.example.com/save" || got["url"] != "https://blog.example.com/p/" {
		t.Errorf("got %v, want endpoint+url", got)
	}
	if len(got) != 2 {
		t.Errorf("request JSON has %d keys, want exactly endpoint and url", len(got))
	}
}
