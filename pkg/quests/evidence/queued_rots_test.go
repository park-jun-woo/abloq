//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what queuedRots가 rot_urls JSON을 디코드하고 부재는 nil, 불량 JSON은 에러인지 검증
package evidence

import (
	"reflect"
	"testing"
)

func TestQueuedRots(t *testing.T) {
	got, err := queuedRots(map[string]string{"rot_urls": `["https://gone.example/x"]`})
	if err != nil || !reflect.DeepEqual(got, []string{"https://gone.example/x"}) {
		t.Errorf("got = %v (%v)", got, err)
	}
	got, err = queuedRots(map[string]string{})
	if err != nil || got != nil {
		t.Errorf("absent rot_urls must yield nil: %v (%v)", got, err)
	}
	if _, err := queuedRots(map[string]string{"rot_urls": "not json"}); err == nil {
		t.Error("malformed rot_urls: want error")
	}
}
