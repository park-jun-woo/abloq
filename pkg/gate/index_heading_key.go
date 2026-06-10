//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what 헤딩 키 1개의 언어별 현지화 헤딩을 정규화해 역색인(언어→정규화 헤딩→키)에 등록
package gate

// indexHeadingKey registers one heading key's localized texts into the
// per-language reverse index.
func indexHeadingKey(byLang map[string]map[string]string, key string, langs map[string]string) {
	for lang, text := range langs {
		if byLang[lang] == nil {
			byLang[lang] = map[string]string{}
		}
		byLang[lang][normText(text)] = key
	}
}
