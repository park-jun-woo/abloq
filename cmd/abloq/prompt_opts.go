//ff:func feature=init type=command control=sequence
//ff:what 대화형 init — 제목/baseURL/저자/언어/섹션 5문항, 빈 입력은 플래그·기본값 유지(비대화형 폴백과 동일 값)
package main

import (
	"bufio"
	"io"
	"strings"
)

// promptOpts asks the five init questions; empty answers keep the defaults,
// so EOF (non-tty stdin) degrades to the non-interactive behavior.
func promptOpts(out io.Writer, in io.Reader, o initOpts) initOpts {
	sc := bufio.NewScanner(in)
	o.Title = promptLine(out, sc, "site title", o.Title)
	o.BaseURL = promptLine(out, sc, "base URL", o.BaseURL)
	o.Author = promptLine(out, sc, "author", o.Author)
	langs := promptLine(out, sc, "languages (comma-separated, first = default)", strings.Join(o.Languages, ","))
	o.Languages = strings.Split(langs, ",")
	sections := promptLine(out, sc, "sections (comma-separated)", strings.Join(o.Sections, ","))
	o.Sections = strings.Split(sections, ",")
	return o
}
