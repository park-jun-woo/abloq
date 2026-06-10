//ff:func feature=init type=command control=sequence
//ff:what promptLine이 답변을 트림해 반환하고 빈 줄과 EOF에서 기본값으로 폴백하는지 검증
package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestPromptLine(t *testing.T) {
	var out bytes.Buffer
	sc := bufio.NewScanner(strings.NewReader("  answer  \n\n"))
	if got := promptLine(&out, sc, "q1", "def"); got != "answer" {
		t.Errorf("answer = %q, want %q", got, "answer")
	}
	if got := promptLine(&out, sc, "q2", "def"); got != "def" {
		t.Errorf("empty line = %q, want default", got)
	}
	if got := promptLine(&out, sc, "q3", "def"); got != "def" {
		t.Errorf("EOF = %q, want default", got)
	}
	if !strings.Contains(out.String(), "q1 [def]: ") {
		t.Errorf("prompt label missing: %q", out.String())
	}
}
