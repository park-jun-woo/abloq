//ff:type feature=gen type=schema
//ff:what 파생물 1개 — 블로그 루트 기준 상대 경로와 생성 바이트, generate/check의 공통 모델
package gen

// Output is one derived file: path relative to the blog root plus its exact bytes.
type Output struct {
	Path string
	Data []byte
}
