//ff:func feature=init type=command control=sequence
//ff:what 프롬프트 1문항 — 라벨과 기본값 출력 후 한 줄 입력, 빈 입력·EOF면 기본값 반환
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// promptLine reads one answer; empty or EOF keeps def.
func promptLine(out io.Writer, sc *bufio.Scanner, label, def string) string {
	fmt.Fprintf(out, "%s [%s]: ", label, def)
	if !sc.Scan() {
		fmt.Fprintln(out)
		return def
	}
	ans := strings.TrimSpace(sc.Text())
	if ans == "" {
		return def
	}
	return ans
}
