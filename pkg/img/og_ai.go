//ff:func feature=image type=generator control=iteration dimension=1
//ff:what AI OG 본체 — 주입된 (variant, Provider) 쌍 목록을 안×count 순차 실행, 성공·실패를 건별 OGOutcome으로 집계
//ff:why 부분 실패가 성공분을 지우지 않도록 에러를 모아 반환 — exit 코드 판단(부분 실패=1)은 cmd 계층 몫 (BUG002)
package img

import "context"

// OGAI executes every variant Count times and returns one outcome per
// attempt, successes and failures alike. local never comes through here —
// the deterministic card goes straight to RenderOG (see og.go).
func OGAI(ctx context.Context, spec OGAISpec, variants []OGVariant) []OGOutcome {
	var outcomes []OGOutcome
	for _, job := range ogAIJobs(spec, variants) {
		err := ogGenerateOne(ctx, spec, job)
		outcomes = append(outcomes, OGOutcome{
			Variant: job.variant.Name, N: job.n, Model: job.variant.Model,
			Path: job.out, Err: err,
		})
	}
	return outcomes
}
