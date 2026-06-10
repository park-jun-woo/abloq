//ff:type feature=visibility type=client topic=crawl
//ff:what 로컬 디렉토리 로그 소스 — 파일명이 곧 키, CLI 단발 분석·Hurl 픽스처·골든 테스트용
package cflog

// DirSource serves log objects from a local directory: every regular file
// name is a key. Tests, Hurl fixtures and the CLI one-shot analysis use it;
// production S3 access is S3Source.
type DirSource struct {
	Root string
}
