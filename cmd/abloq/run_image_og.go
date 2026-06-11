//ff:func feature=cli type=command control=sequence
//ff:what image og 실행 본체 — provider 해석(플래그 > blog.yaml image.og > local), local은 RenderOG 직행, AI는 안 해석→건수 echo→안×count 실행→결과·채택 안내
//ff:why AI 경로는 generate/check 어디에도 결선되지 않는 명시 호출 1회 — 후보는 검토 후 mv로 채택하는 1회성 자산이다 (BUG002)
package main

import (
	"context"
	"fmt"
	"io"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// ogDraftDir is where multi-candidate runs land: outside every build/derived
// path, so Hugo and generate/check never pick drafts up.
const ogDraftDir = "files/og"

// runImageOG resolves the provider (flag > blog.yaml image.og > local) and
// dispatches: local goes straight to the deterministic card (pre-Phase022
// bytes), AI providers run the injected (variant, Provider) pairs and report
// candidates for review.
func runImageOG(out io.Writer, opts imageOGOpts) error {
	if opts.Count < 1 {
		return fmt.Errorf("--count must be at least 1, got %d", opts.Count)
	}
	cfg, err := loadImageOG(out, ".")
	if err != nil {
		return err
	}
	provider := cfg.OGProvider()
	if opts.ProviderSet && opts.Provider != "" {
		provider = opts.Provider
	}
	multi := opts.AllVariants || opts.VariantList != "" || opts.Count > 1
	if provider == "local" {
		if multi {
			return fmt.Errorf("provider local is deterministic — --variant/--all-variants/--count need an AI provider (blog.yaml image.og or --provider)")
		}
		return runImageOGLocal(out, opts)
	}
	specs, err := resolveOGVariants(cfg.OG, opts)
	if err != nil {
		return err
	}
	runs, err := buildOGRuns(provider, specs, opts)
	if err != nil {
		return err
	}
	printOGPlan(out, runs, opts.Count)
	spec := img.OGAISpec{
		Slug: opts.Slug, Title: opts.Title, Brand: opts.Brand, FontPath: opts.FontPath,
		OutDir: opts.OutDir, DraftDir: ogDraftDir, Multi: multi, Count: opts.Count,
	}
	return printOGOutcomes(out, opts, img.OGAI(context.Background(), spec, runs), multi)
}
