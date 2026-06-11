//ff:func feature=cli type=command control=selection
//ff:what 생성할 안 목록 결정 — --all-variants는 선언 전체, --variant는 지정 목록(안 선언이 플래그보다 우선), 무지정은 기본 설정 1안(플래그가 blog.yaml 오버라이드)
package main

import (
	"errors"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

// resolveOGVariants picks the candidate set. With --variant/--all-variants
// the variant declarations win over --model/--overlay; on the single default
// path explicit flags override the blog.yaml site-wide values.
func resolveOGVariants(og blogyaml.ImageOG, opts imageOGOpts) ([]blogyaml.OGVariantSpec, error) {
	switch {
	case opts.AllVariants:
		specs := og.ResolvedVariants()
		if len(specs) == 0 {
			return nil, errors.New("--all-variants: blog.yaml image.og declares no variants")
		}
		return specs, nil
	case opts.VariantList != "":
		return ogVariantsFromList(og, opts.VariantList)
	}
	d := og.DefaultVariant()
	if opts.ModelSet {
		d.Model = opts.Model
	}
	if opts.OverlaySet {
		d.Overlay = opts.Overlay
	}
	return []blogyaml.OGVariantSpec{d}, nil
}
