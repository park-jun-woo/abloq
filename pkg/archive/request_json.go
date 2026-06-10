//ff:func feature=archive type=client control=sequence
//ff:what 영수증 request JSON 조립 — {"endpoint", "url"} (자격증명은 절대 기록하지 않는다)
package archive

import "encoding/json"

// requestJSON records what was submitted where. Credentials (IndexNow key,
// GSC token, Wayback keys) never enter the receipt.
func requestJSON(endpoint, target string) json.RawMessage {
	data, err := json.Marshal(map[string]string{"endpoint": endpoint, "url": target})
	if err != nil {
		return json.RawMessage(`{}`)
	}
	return data
}
