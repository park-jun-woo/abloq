//ff:func feature=cli type=command control=iteration dimension=1
//ff:what 확정 안 목록에 Provider 인스턴스와 치환 완료 프롬프트를 짝지어 실행용 (variant, Provider) 쌍으로 변환 — 키 부재는 여기서 즉시 에러
package main

import (
	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/img"
)

// buildOGRuns turns normalized variant specs into injected execution pairs:
// each gets its resolved provider instance and its rendered prompt
// ({title}/{summary}/{brand} substituted; summary is resolved from the
// article's front matter, empty on the local/no-blog.yaml paths).
func buildOGRuns(provider string, specs []blogyaml.OGVariantSpec, opts imageOGOpts) ([]img.OGVariant, error) {
	var runs []img.OGVariant
	for _, s := range specs {
		p, model, err := resolveOGProvider(provider, s.Model)
		if err != nil {
			return nil, err
		}
		runs = append(runs, img.OGVariant{
			Name:     s.Name,
			Model:    model,
			Overlay:  s.Overlay,
			Prompt:   blogyaml.OGPrompt(s.Prompt, opts.Title, opts.Summary, opts.Brand),
			Provider: p,
		})
	}
	return runs, nil
}
