//ff:func feature=queueio type=parser control=iteration dimension=1
//ff:what 디스크 큐 파일(Serialize 산출형) → Item — 라인 단위 strict 파싱, Serialize와 왕복 동등 (Phase018 소비 퀘스트 Seed가 사용)
//ff:why DecodeRows는 jsonb_agg JSON 파서라 디스크 YAML형을 못 읽는다 — 소비 퀘스트의 Seed는 파일이 유일한 입력이므로 직렬화의 정확한 역연산이 필요하다
package queueio

import "strings"

// Deserialize parses one queue file written by Serialize back into its Item.
// The parse is strict (unknown or malformed lines error) so a hand-tampered
// queue file fails loudly at Seed instead of seeding a half-read item.
// Deserialize(Serialize(it)) == it for every valid item.
func Deserialize(data []byte) (Item, error) {
	it := Item{Payload: map[string]string{}}
	for _, ln := range strings.Split(string(data), "\n") {
		if err := applyQueueLine(&it, ln); err != nil {
			return Item{}, err
		}
	}
	return it, nil
}
