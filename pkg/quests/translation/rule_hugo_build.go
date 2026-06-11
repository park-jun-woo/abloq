//ff:func feature=quest type=rule control=sequence
//ff:what [hugo-build] 인스턴스 전체 hugo 빌드 0 에러 검사 — 임시 출력 디렉토리로 빌드, 실패 출력은 Fact로
//ff:why hugo 부재는 이 룰의 SKIP이 아니라 Prepare 에러(제출 중단·try 미소진) — SKIP 진단은 reins 계약상 침묵 통과 = PATH 조작 치즈 구멍. 빌드 산출물 검사(hreflang-complete)는 repo/CI 레벨 abloq gate 몫이고 퀘스트는 빌드 성공만 잠근다 (Phase017 계획)
package translation

import (
	"os"
	"os/exec"

	rgate "github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleHugoBuild builds the whole instance with hugo into a throwaway
// destination and fires on any non-zero exit, quoting the build output.
func ruleHugoBuild() rgate.Rule {
	return rgate.Rule{
		Meta: rgate.RuleMeta{ID: "hugo-build", Level: rgate.LevelFail,
			Desc: "hugo builds the whole instance with zero errors"},
		Check: func(ctx rgate.Context) (bool, quest.Fact) {
			sub := ctx.Submission.(*Submission)
			bin, err := hugoPath()
			if err != nil {
				return true, quest.Fact{Where: sub.Root, Expected: "hugo binary in PATH", Actual: err.Error()}
			}
			tmp, err := os.MkdirTemp("", "abloq-hugo-*")
			if err != nil {
				return true, quest.Fact{Where: sub.Root, Expected: "writable temp build dir", Actual: err.Error()}
			}
			defer os.RemoveAll(tmp)
			out, err := exec.Command(bin, "--quiet", "--noBuildLock", "-s", sub.Root, "-d", tmp).CombinedOutput()
			if err == nil {
				return false, quest.Fact{}
			}
			return true, quest.Fact{Where: sub.Root,
				Expected: "hugo build with zero errors",
				Actual:   truncOut(out, err)}
		},
	}
}
