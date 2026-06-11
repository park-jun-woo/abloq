//ff:func feature=queueio type=parser control=sequence
//ff:what Deserialize가 Serialize 산출과 왕복 동등(keys·payload·빈 payload 포함)하고 변조 라인은 에러인지 검증
package queueio

import (
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	it := Item{
		Kind: "evidence", Slug: "post-a", Lang: "ko", Section: "tech",
		Priority: 3,
		Keys:     []string{"ko/tech/post-a", "en/tech/post-a"},
		Payload: map[string]string{
			"claims":   `[{"hash":"abcd","loc":"content/ko/tech/post-a.md:5","text":"x 40% 증가"}]`,
			"rot_urls": `["https://gone.example/x"]`,
		},
	}
	got, err := Deserialize(Serialize(it))
	if err != nil {
		t.Fatalf("Deserialize: %v", err)
	}
	if !reflect.DeepEqual(got, it) {
		t.Errorf("round trip mismatch:\n got %+v\nwant %+v", got, it)
	}
	empty := Item{Kind: "refresh", Slug: "s", Lang: "ko", Section: "tech", Payload: map[string]string{}}
	got, err = Deserialize(Serialize(empty))
	if err != nil {
		t.Fatalf("Deserialize empty payload: %v", err)
	}
	if !reflect.DeepEqual(got, empty) {
		t.Errorf("empty-payload round trip mismatch:\n got %+v\nwant %+v", got, empty)
	}
	if _, err := Deserialize([]byte("kind: \"x\"\ngarbage line\n")); err == nil {
		t.Error("unrecognized line must error")
	}
	if _, err := Deserialize([]byte("priority: NaN\n")); err == nil {
		t.Error("malformed priority must error")
	}
	if _, err := Deserialize([]byte("payload:\n  k: unquoted\n")); err == nil {
		t.Error("unquoted payload value must error")
	}
}
