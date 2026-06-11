//ff:type feature=quest type=schema topic=queue
//ff:what 제출물 JSON(submit --in) — 루트 기준 대상 글 경로 1개 (후보 글 변경은 변경 집합으로 검사, 별도 신고 불요)
package cluster

// SubmitFile is the decoded submission JSON: the seeded target article path
// (relative to the instance root). Candidate-article edits need no
// declaration — queue-scope checks the whole working-tree change set against
// the allowed set, and the re-scan walks the repository anyway.
type SubmitFile struct {
	Article string `json:"article"`
}
