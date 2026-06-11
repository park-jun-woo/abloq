//ff:func feature=queueio type=parser control=iteration dimension=1
//ff:what DB payload에서 keys 키(JSON 문자열)를 []string 1급 값으로 분리하고 사본에서 제거 (EncodeRows의 역연산)
package queueio

import "encoding/json"

// liftKeys splits the DB payload into the decoded per-language key list and
// a copy of the remaining keys. A payload without the keys entry (legacy
// rows) yields nil — the queue file simply carries no keys: block.
func liftKeys(payload map[string]string) ([]string, map[string]string) {
	raw, ok := payload["keys"]
	if !ok {
		return nil, payload
	}
	rest := make(map[string]string, len(payload))
	for k, v := range payload {
		rest[k] = v
	}
	delete(rest, "keys")
	var keys []string
	if err := json.Unmarshal([]byte(raw), &keys); err != nil {
		return nil, rest
	}
	return keys, rest
}
