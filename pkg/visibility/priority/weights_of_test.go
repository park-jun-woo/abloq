//ff:func feature=visibility type=scorer control=sequence
//ff:what WeightsOf가 blog.yaml priority_weights를 계수 그대로 옮기는지 검증
package priority

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestWeightsOf(t *testing.T) {
	w := WeightsOf(blogyaml.PriorityWeights{Fetcher: 3, Train: 1, GSC: 4, Citation: 2})
	want := Weights{Fetcher: 3, Train: 1, GSC: 4, Citation: 2}
	if w != want {
		t.Errorf("want %+v, got %+v", want, w)
	}
}
