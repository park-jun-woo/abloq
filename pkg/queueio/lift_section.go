//ff:func feature=queueio type=parser control=iteration dimension=1
//ff:what DB payload에서 section 키를 1급 값으로 분리하고 나머지를 사본으로 반환 (EncodeRows의 역연산)
package queueio

// liftSection splits the DB payload into the section value and a copy of the
// remaining keys.
func liftSection(payload map[string]string) (string, map[string]string) {
	section := ""
	rest := make(map[string]string, len(payload))
	for k, v := range payload {
		if k == "section" {
			section = v
			continue
		}
		rest[k] = v
	}
	return section, rest
}
