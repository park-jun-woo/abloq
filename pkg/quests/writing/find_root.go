//ff:func feature=quest type=parser control=sequence
//ff:what 디렉토리에서 위로 blog.yaml을 탐색해 인스턴스 루트 반환 — 공용 추출본(quests/common.FindRoot) 위임
//ff:why Phase017에서 번역 퀘스트와 공유하도록 구현을 pkg/quests/common으로 추출 — 복제 금지, writing은 추출본을 쓴다
package writing

import "github.com/park-jun-woo/abloq/pkg/quests/common"

// findRoot walks up from dir to the nearest directory containing blog.yaml.
// The implementation lives in pkg/quests/common (Phase017 extraction).
func findRoot(dir string) (string, error) {
	return common.FindRoot(dir)
}
