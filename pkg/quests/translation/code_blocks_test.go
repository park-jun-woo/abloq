//ff:func feature=quest type=parser control=sequence
//ff:what codeBlocks 검증 — 펜스 블록을 여는 펜스(언어 태그)+내용으로 추출, 블록 수·내용 보존
package translation

import (
	"strings"
	"testing"
)

func TestCodeBlocks(t *testing.T) {
	md := "---\ntitle: x\n---\n\n```go\na := 1\n```\n\ntext\n\n```sh\nrun\n\nmore\n```\n"
	blocks := codeBlocks(docOf(t, "en", md))
	if len(blocks) != 2 {
		t.Fatalf("blocks = %d, want 2", len(blocks))
	}
	if blocks[0] != "```go\na := 1" {
		t.Errorf("blocks[0] = %q", blocks[0])
	}
	if !strings.Contains(blocks[1], "run\n\nmore") {
		t.Errorf("blocks[1] = %q, want blank line kept inside", blocks[1])
	}
}
