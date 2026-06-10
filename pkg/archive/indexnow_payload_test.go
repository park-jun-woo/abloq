//ff:func feature=archive type=client control=sequence
//ff:what indexNowPayload가 host·key·규약 keyLocation·urlList를 채우고 INDEXNOW_KEY_LOCATION 오버라이드와 host 없는 target 거부를 지키는지 검증
package archive

import "testing"

func TestIndexNowPayload(t *testing.T) {
	pending := []Pending{
		{Kind: KindIndexNow, Target: "https://blog.example.com/a/"},
		{Kind: KindIndexNow, Target: "https://blog.example.com/b/"},
	}
	payload, err := indexNowPayload("k123", pending)
	if err != nil {
		t.Fatalf("indexNowPayload: %v", err)
	}
	if payload["host"] != "blog.example.com" || payload["key"] != "k123" {
		t.Errorf("host/key = %v/%v", payload["host"], payload["key"])
	}
	if payload["keyLocation"] != "https://blog.example.com/k123.txt" {
		t.Errorf("keyLocation = %v, want protocol default", payload["keyLocation"])
	}
	if urls := payload["urlList"].([]string); len(urls) != 2 {
		t.Errorf("urlList = %v, want both targets", urls)
	}

	t.Setenv("INDEXNOW_KEY_LOCATION", "https://cdn.example.com/key.txt")
	payload, err = indexNowPayload("k123", pending)
	if err != nil || payload["keyLocation"] != "https://cdn.example.com/key.txt" {
		t.Errorf("keyLocation override = %v (err=%v)", payload["keyLocation"], err)
	}

	if _, err := indexNowPayload("k123", []Pending{{Target: "not-a-url"}}); err == nil {
		t.Error("host-less target must fail")
	}
}
