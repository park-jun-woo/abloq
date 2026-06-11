//ff:func feature=quest type=parser control=sequence
//ff:what 루트 기준 글 경로에서 Item Key 부품(lang/section/slug) 파생 — 공용 추출본(quests/common.KeyParts) 위임
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// keyParts derives the item key parts from a root-relative article path.
// The implementation lives in pkg/quests/common (Phase017 extraction).
func keyParts(article string) (lang, section, slug string, ok bool) {
	return common.KeyParts(article)
}
