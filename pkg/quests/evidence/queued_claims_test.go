//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what queuedClaims가 claims JSON을 해시 집합으로 디코드하고 부재는 빈 집합, 불량 JSON은 에러인지 검증
package evidence

import "testing"

func TestQueuedClaims(t *testing.T) {
	got, err := queuedClaims(map[string]string{
		"claims": `[{"hash":"aaaa","loc":"a.md:1","text":"x"},{"hash":"bbbb","loc":"a.md:2","text":"y"}]`})
	if err != nil || len(got) != 2 || !got["aaaa"] || !got["bbbb"] {
		t.Errorf("got = %v (%v)", got, err)
	}
	got, err = queuedClaims(map[string]string{})
	if err != nil || len(got) != 0 || got == nil {
		t.Errorf("absent claims must yield an empty set: %v (%v)", got, err)
	}
	if _, err := queuedClaims(map[string]string{"claims": "not json"}); err == nil {
		t.Error("malformed claims: want error")
	}
}
