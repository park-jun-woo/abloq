//ff:func feature=cli type=output control=sequence
//ff:what printEvidenceQueue가 항목마다 큐 파일 경로와 priority 한 줄을 내는지 검증
package main

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

func TestPrintEvidenceQueue(t *testing.T) {
	var out bytes.Buffer
	printEvidenceQueue(&out, []queueio.Item{
		{Kind: "evidence", Lang: "ko", Section: "tech", Slug: "post-a", Priority: 2},
	})
	want := filepath.Join("quests", "queue", "evidence-ko-tech-post-a.yaml") + "\tpriority=2\n"
	if out.String() != want {
		t.Errorf("output = %q, want %q", out.String(), want)
	}
	out.Reset()
	printEvidenceQueue(&out, nil)
	if out.String() != "" {
		t.Errorf("no items must print nothing: %q", out.String())
	}
}
