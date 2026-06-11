//ff:func feature=image type=generator control=iteration dimension=1
//ff:what 안 목록 × 안당 샘플 수를 잡 목록으로 전개 — 경로는 ogJobOut이 직행/드래프트로 분기
package img

// ogAIJobs expands variants × spec.Count into flat jobs with their output
// paths, so the execution loop stays one-dimensional.
func ogAIJobs(spec OGAISpec, variants []OGVariant) []ogJob {
	count := spec.Count
	if count < 1 {
		count = 1
	}
	var jobs []ogJob
	for _, v := range variants {
		for n := 1; n <= count; n++ {
			jobs = append(jobs, ogJob{variant: v, n: n, out: ogJobOut(spec, v.Name, n)})
		}
	}
	return jobs
}
