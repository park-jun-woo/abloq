//ff:func feature=scan type=client control=sequence topic=evidence
//ff:what mergeResults가 채널의 부분 결과 맵 n개를 하나로 합치는지 검증
package evidence

import "testing"

func TestMergeResults(t *testing.T) {
	out := make(chan map[string]string, 2)
	out <- map[string]string{"https://a.example/x": "ok"}
	out <- map[string]string{"https://b.example/y": "hard"}
	res := mergeResults(out, 2)
	if len(res) != 2 || res["https://a.example/x"] != "ok" || res["https://b.example/y"] != "hard" {
		t.Errorf("mergeResults = %v", res)
	}
}
