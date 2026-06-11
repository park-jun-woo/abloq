//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runImageOG AI 경로 통합 검증 — 단일 직행/다중 안 드래프트/건수 echo/모델 echo/mv 채택 안내/부분 실패 exit/해석 순서/키 부재를 로컬 스텁 서버로 (실 네트워크 0)
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunImageOGAI(t *testing.T) {
	dir := writeOGBlogFixture(t)
	chdirTemp(t, dir)
	srv := serveGeminiStub(t)
	t.Setenv("GEMINI_API_KEY", "test-key")
	t.Setenv("GEMINI_API_BASE", srv.URL)

	// single default call: blog.yaml provider gemini -> direct adopted path
	var out bytes.Buffer
	opts := imageOGOpts{Slug: "post", Title: "My Post", OutDir: "static/images", Count: 1}
	if err := runImageOG(&out, opts); err != nil {
		t.Fatalf("single AI call: %v\n%s", err, out.String())
	}
	if _, err := os.Stat(filepath.Join("static", "images", "post.webp")); err != nil {
		t.Errorf("direct output missing: %v", err)
	}
	for _, want := range []string{"planned: 1 image(s)", "(model good-model)", `image: "/images/post.webp"`} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("single output missing %q:\n%s", want, out.String())
		}
	}

	// multi: --variant minimal,photo --count 2 -> 4 drafts + adoption hint
	out.Reset()
	opts = imageOGOpts{Slug: "post", Title: "My Post", OutDir: "static/images", VariantList: "minimal,photo", Count: 2}
	if err := runImageOG(&out, opts); err != nil {
		t.Fatalf("multi: %v\n%s", err, out.String())
	}
	for _, name := range []string{"minimal-1.webp", "minimal-2.webp", "photo-1.webp", "photo-2.webp"} {
		if _, err := os.Stat(filepath.Join("files", "og", "post", name)); err != nil {
			t.Errorf("draft %s missing: %v", name, err)
		}
	}
	for _, want := range []string{"planned: 4 image(s)", "(model photo-model)", "mv " + filepath.Join("files", "og", "post", "minimal-1.webp")} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("multi output missing %q:\n%s", want, out.String())
		}
	}

	// partial failure: bad variant fails, sibling survives, error returned
	out.Reset()
	opts = imageOGOpts{Slug: "mix", Title: "T", OutDir: "static/images", VariantList: "minimal,bad", Count: 1}
	err := runImageOG(&out, opts)
	if err == nil || !strings.Contains(err.Error(), "1 of 2") {
		t.Fatalf("partial failure: want '1 of 2' error, got %v", err)
	}
	if _, err := os.Stat(filepath.Join("files", "og", "mix", "minimal-1.webp")); err != nil {
		t.Errorf("successful sibling lost: %v", err)
	}
	if !strings.Contains(out.String(), "failed: bad-1") {
		t.Errorf("failure detail missing:\n%s", out.String())
	}

	// unknown variant name
	out.Reset()
	opts = imageOGOpts{Slug: "post", Title: "T", OutDir: "static/images", VariantList: "nope", Count: 1}
	if err := runImageOG(&out, opts); err == nil || !strings.Contains(err.Error(), "nope") {
		t.Errorf("unknown variant: want error naming it, got %v", err)
	}

	// resolution order: --provider local beats blog.yaml gemini (no API call)
	out.Reset()
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("GOOGLE_API_KEY", "")
	opts = imageOGOpts{Slug: "loc", Title: "T", OutDir: "static/images", Provider: "local", ProviderSet: true, Count: 1}
	if err := runImageOG(&out, opts); err != nil {
		t.Fatalf("--provider local override: %v", err)
	}
	if _, err := os.Stat(filepath.Join("static", "images", "loc.webp")); err != nil {
		t.Errorf("local override output missing: %v", err)
	}

	// key absence: blog.yaml gemini + no key -> clear exit-1 error
	out.Reset()
	opts = imageOGOpts{Slug: "post", Title: "T", OutDir: "static/images", Count: 1}
	if err := runImageOG(&out, opts); err == nil || !strings.Contains(err.Error(), "GEMINI_API_KEY") {
		t.Errorf("missing key: want clear diagnosis, got %v", err)
	}

	// local + multi flags is a contradiction
	out.Reset()
	opts = imageOGOpts{Slug: "post", Title: "T", OutDir: "static/images", Provider: "local", ProviderSet: true, Count: 3}
	if err := runImageOG(&out, opts); err == nil || !strings.Contains(err.Error(), "deterministic") {
		t.Errorf("local+count: want contradiction error, got %v", err)
	}

	// count < 1
	opts = imageOGOpts{Slug: "post", Title: "T", OutDir: "static/images", Count: 0}
	if err := runImageOG(&out, opts); err == nil || !strings.Contains(err.Error(), "--count") {
		t.Errorf("count 0: want error, got %v", err)
	}
}
