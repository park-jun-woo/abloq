//ff:func feature=cli type=output control=iteration dimension=1
//ff:what 생성 결과 출력 — 성공 경로(+모델 echo)·실패 내역, 다중 안은 후보 목록과 mv 채택 안내, 부분 실패는 성공분 보존 후 에러(exit 1)
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/img"
)

// printOGOutcomes reports every attempt. Multi-candidate runs end with the
// adoption hint (candidates are already final-format WebP — one mv adopts).
// Any failure returns an error so the pipeline exits 1 without discarding
// the successful candidates.
func printOGOutcomes(out io.Writer, opts imageOGOpts, outcomes []img.OGOutcome, multi bool) error {
	var ok []img.OGOutcome
	for _, o := range outcomes {
		if o.Err != nil {
			fmt.Fprintf(out, "failed: %s-%d (model %s): %v\n", o.Variant, o.N, o.Model, o.Err)
			continue
		}
		fmt.Fprintf(out, "%s (model %s)\n", o.Path, o.Model)
		ok = append(ok, o)
	}
	if !multi && len(ok) == 1 {
		fmt.Fprintf(out, "  front matter: image: \"/images/%s.webp\"\n", opts.Slug)
	}
	if multi && len(ok) > 0 {
		fmt.Fprintf(out, "review the %d candidate(s), then adopt one (already 1200x630 WebP):\n", len(ok))
		fmt.Fprintf(out, "  mv %s %s\n", ok[0].Path, filepath.Join(opts.OutDir, opts.Slug+".webp"))
	}
	if failed := len(outcomes) - len(ok); failed > 0 {
		return fmt.Errorf("%d of %d generation(s) failed — successful candidates are kept", failed, len(outcomes))
	}
	return nil
}
