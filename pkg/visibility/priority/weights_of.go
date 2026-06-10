//ff:func feature=visibility type=scorer control=sequence
//ff:what blog.yaml geo.priority_weights → Weights 변환 — 주입 지점(백엔드 func·CLI)이 공유하는 단일 매핑
package priority

import "github.com/park-jun-woo/abloq/pkg/blogyaml"

// WeightsOf converts the blog.yaml priority_weights section into the scorer
// coefficients. The mapping is the single conversion point shared by the
// backend funcs and the CLI, so both inject identical weights.
func WeightsOf(w blogyaml.PriorityWeights) Weights {
	return Weights{
		Fetcher:  w.Fetcher,
		Train:    w.Train,
		GSC:      w.GSC,
		Citation: w.Citation,
	}
}
