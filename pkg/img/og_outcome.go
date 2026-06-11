//ff:type feature=image type=schema
//ff:what AI OG 생성 1건의 결과 — 안 이름/샘플 번호/모델/기록 경로/실패 에러, 성공·실패 집계의 단위
package img

// OGOutcome reports one generation attempt. Err==nil means Path holds the
// written WebP; on failure Path is where it would have been written. Partial
// failures keep their successful siblings — the caller decides the exit code.
type OGOutcome struct {
	Variant string
	N       int
	Model   string
	Path    string
	Err     error
}
