//ff:type feature=quest type=schema
//ff:what 제출물 JSON(submit --in) — 루트 기준 글/작업 로그/REVIEW 기록 경로 3개
package writing

// SubmitFile is the decoded submission JSON: the three artifact paths
// (relative to the instance root) the gate inspects. Article must equal the
// seeded target path — submitting another file is rejected in Prepare.
type SubmitFile struct {
	Article string `json:"article"`
	Worklog string `json:"worklog"`
	Review  string `json:"review"`
}
