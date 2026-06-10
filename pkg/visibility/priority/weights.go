//ff:type feature=visibility type=schema
//ff:what 측정 신호 가중 계수 — fetcher/train/gsc/citation, blog.yaml geo.priority_weights에서 주입 (Phase014)
package priority

// Weights are the measurement-signal coefficients of the Measured scorer.
// They mirror blog.yaml geo.priority_weights (WeightsOf converts); the
// fetcher weight is the highest by default — a user-triggered fetch is the
// strongest consumption evidence (§6.1).
type Weights struct {
	Fetcher  int64
	Train    int64
	GSC      int64
	Citation int64
}
