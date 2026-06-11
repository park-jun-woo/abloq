//ff:func feature=cli type=command control=sequence
//ff:what insight match 에러 경로 검증 — 명세 없음, 글 없음, 스키마 에러, 섹션 불일치
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunInsightMatchError(t *testing.T) {
	var out bytes.Buffer
	if err := runInsightMatch(&out, filepath.Join(t.TempDir(), "absent.insight.yaml"), "x.md"); err == nil {
		t.Errorf("want IO error for missing insight file, got nil")
	}
	insightPath, articlePath := writeInsightFixture(t,
		"topic: t\nsection: tech\nclaims:\n  - id: a\n    text: x\n    kind: claim\n    anchors: [\"ratchet\"]\n")
	if err := runInsightMatch(&out, insightPath, filepath.Join(t.TempDir(), "absent.md")); err == nil {
		t.Errorf("want IO error for missing article, got nil")
	}
	if err := os.WriteFile(insightPath, []byte("topic: t\nsection: tech\nclaims: []\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if err := runInsightMatch(&out, insightPath, articlePath); err == nil || !strings.Contains(out.String(), "insight-claims-min") {
		t.Errorf("want claims-min validation error, got err=%v output=%q", err, out.String())
	}
	if err := os.WriteFile(insightPath, []byte("topic: t\nsection: opinion\nclaims:\n  - id: a\n    text: x\n    kind: claim\n    anchors: [\"ratchet\"]\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	err := runInsightMatch(&out, insightPath, articlePath)
	if err == nil || !strings.Contains(err.Error(), "section mismatch") {
		t.Errorf("want section mismatch error, got %v", err)
	}
}
