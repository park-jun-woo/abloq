//ff:func feature=cli type=output control=sequence
//ff:what printOGOutcomes 검증 — 단일 성공의 front matter 안내, 다중 안의 mv 채택 안내, 부분 실패 시 성공분 보존 에러
package main

import (
	"bytes"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/img"
)

func TestPrintOGOutcomes(t *testing.T) {
	opts := imageOGOpts{Slug: "post", OutDir: "static/images"}

	// single success: path + model echo + front matter hint, nil error
	var out bytes.Buffer
	ok := img.OGOutcome{Variant: "default", N: 1, Model: "m1", Path: "static/images/post.webp"}
	if err := printOGOutcomes(&out, opts, []img.OGOutcome{ok}, false); err != nil {
		t.Fatalf("single success: %v", err)
	}
	if !strings.Contains(out.String(), "static/images/post.webp (model m1)") ||
		!strings.Contains(out.String(), `front matter: image: "/images/post.webp"`) {
		t.Errorf("single success output = %q", out.String())
	}

	// multi: adoption hint with mv from first success to final path; failure
	// line reported, error returned while successes are kept
	out.Reset()
	fail := img.OGOutcome{Variant: "bold", N: 2, Model: "m2", Path: "x", Err: errors.New("boom")}
	err := printOGOutcomes(&out, opts, []img.OGOutcome{fail, ok}, true)
	if err == nil || !strings.Contains(err.Error(), "1 of 2") {
		t.Fatalf("multi partial failure: err = %v, want 1 of 2", err)
	}
	s := out.String()
	if !strings.Contains(s, "failed: bold-2 (model m2): boom") {
		t.Errorf("missing failure line: %q", s)
	}
	wantMv := "mv static/images/post.webp " + filepath.Join("static/images", "post.webp")
	if !strings.Contains(s, "review the 1 candidate(s)") || !strings.Contains(s, wantMv) {
		t.Errorf("missing adoption hint: %q", s)
	}
	if strings.Contains(s, "front matter") {
		t.Errorf("multi run must not print the front matter hint: %q", s)
	}
}
