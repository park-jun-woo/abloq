//ff:type feature=quest type=schema topic=queue
//ff:what 제출물 JSON(submit --in) — 루트 기준 갱신 글 경로 1개 (이 퀘스트의 산출물은 갱신된 글뿐)
package refresh

// SubmitFile is the decoded submission JSON: the single artifact path
// (relative to the instance root) the gate inspects. Article must equal the
// seeded target path — submitting another file is rejected in Prepare.
type SubmitFile struct {
	Article string `json:"article"`
}
