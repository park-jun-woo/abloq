//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what git show HEAD:./경로로 글의 HEAD 스냅샷을 읽음 — HEAD에 없으면(글이 미커밋) 에러, Base nil 폴백 금지 (퀘스트 공용)
//ff:why 글이 HEAD에 없는데 Base nil로 폴백하면 기준선 룰 전체가 침묵 통과한다(치즈) — 큐 소비는 기존 글 수정이 본질이라 HEAD 부재는 인스턴스 상태 오류이고 Prepare 에러(try 미소진)로 중단한다 (2차 검수 확정)
package common

import (
	"fmt"
	"os/exec"
)

// gitHeadRaw reads the HEAD snapshot of a root-relative path. The article
// must exist at HEAD — the queue quests modify committed articles, so a
// missing snapshot aborts Prepare instead of silently dropping the baseline.
func gitHeadRaw(root, rel string) ([]byte, error) {
	out, err := exec.Command("git", "-C", root, "show", "HEAD:./"+rel).Output()
	if err != nil {
		return nil, fmt.Errorf("git show HEAD:./%s: %w — the target article must be committed before consuming its queue item (no Base-nil fallback)", rel, err)
	}
	return out, nil
}
