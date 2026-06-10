//ff:type feature=init type=schema
//ff:what abloq init 입력 — 제목/baseURL/저자/언어/섹션과 대화형 여부, 비대화형이면 플래그·기본값만 사용
package main

// initOpts carries the answers `abloq init` needs. Interactive=false is the
// agent path: flags and defaults only, no prompts.
type initOpts struct {
	Title       string
	BaseURL     string
	Author      string
	Languages   []string
	Sections    []string
	Interactive bool
}
