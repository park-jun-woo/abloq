//ff:func feature=cli type=command control=sequence
//ff:what resolveOGSummary 검증 — base 언어 유효 slug 매칭 1건/file-stem 미스/미발견 진단/nil-blog 패닉 없이 빈값
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestResolveOGSummary(t *testing.T) {
	dir := writeSummaryFixture(t)
	var out bytes.Buffer
	b, _, err := loadImageOG(&out, dir)
	if err != nil || b == nil {
		t.Fatalf("loadImageOG: blog %v err %v", b, err)
	}

	// match: base language, front matter slug override applied, silent on hit
	out.Reset()
	if got := resolveOGSummary(&out, dir, b, "real-slug"); got != "base summary text" {
		t.Errorf("match by overridden slug = %q, want base summary text", got)
	}
	if out.Len() != 0 {
		t.Errorf("a clean match must be silent, got %q", out.String())
	}

	// the file stem (pre-override) does not match
	out.Reset()
	if got := resolveOGSummary(&out, dir, b, "file-stem"); got != "" {
		t.Errorf("file stem must not match overridden slug, got %q", got)
	}

	// not found: empty + one diagnostic line
	out.Reset()
	if got := resolveOGSummary(&out, dir, b, "nope"); got != "" {
		t.Errorf("missing slug = %q, want empty", got)
	}
	if !strings.Contains(out.String(), "summary 미적용") {
		t.Errorf("not-found must print a diagnostic, got %q", out.String())
	}

	// nil blog (no blog.yaml): skip resolution, empty, NO panic / NO IndexEntries(nil)
	out.Reset()
	if got := resolveOGSummary(&out, dir, nil, "real-slug"); got != "" {
		t.Errorf("nil blog = %q, want empty", got)
	}
	if out.Len() != 0 {
		t.Errorf("nil blog must skip silently, got %q", out.String())
	}
}
