//ff:func feature=cli type=command control=iteration dimension=1
//ff:what insight match 본체 검증 — 미출현 claim은 에러 없이 목록 출력, anchors 빈 claim 경고, 규약 밖 경로 note
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunInsightMatch(t *testing.T) {
	insightPath, articlePath := writeInsightFixture(t,
		"topic: t\nsection: tech\nclaims:\n  - id: hit\n    text: h\n    kind: claim\n    anchors: [\"RATCHET NEVER\"]\n"+
			"  - id: absent\n    text: a\n    kind: prediction\n    anchors: [\"unrelated phrase\"]\n"+
			"  - id: blind\n    text: b\n    kind: claim\n")
	var out bytes.Buffer
	if err := runInsightMatch(&out, insightPath, articlePath); err != nil {
		t.Fatalf("want nil error (missing claims are REVIEW aid), got %v", err)
	}
	got := out.String()
	for _, want := range []string{"insight-claim-anchors-empty", "anchored claims: 1/3", "absent: a", "blind: b"} {
		if !strings.Contains(got, want) {
			t.Errorf("want %q in output, got %q", want, got)
		}
	}
	if strings.Contains(got, "note: insight file is not at the conventional path") {
		t.Errorf("want no convention note for sidecar path, got %q", got)
	}
	other := filepath.Join(t.TempDir(), "elsewhere.insight.yaml")
	data, err := os.ReadFile(insightPath)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(other, data, 0o644); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if err := runInsightMatch(&out, other, articlePath); err != nil {
		t.Fatalf("want nil error from unconventional path, got %v", err)
	}
	if !strings.Contains(out.String(), "note: insight file is not at the conventional path") {
		t.Errorf("want convention note for unconventional path, got %q", out.String())
	}
}
