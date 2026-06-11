//ff:func feature=queueio type=parser control=iteration dimension=1
//ff:what queue_items jsonb_agg JSON → []Row — payload의 section·keys를 1급 필드로 승격, payload 사본에서 제거
package queueio

import "encoding/json"

// DecodeRows parses the jsonb_agg scalar the queue queries return. Section
// and the per-language keys are lifted out of payload into the first-class
// fields (the inverse of EncodeRows); the payload map is copied so neither
// key leaks into the serialized file twice.
func DecodeRows(data []byte) ([]Row, error) {
	var raw []Row
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	rows := make([]Row, 0, len(raw))
	for _, r := range raw {
		r.Section, r.Payload = liftSection(r.Payload)
		r.Keys, r.Payload = liftKeys(r.Payload)
		rows = append(rows, r)
	}
	return rows, nil
}
