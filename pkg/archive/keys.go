//ff:type feature=archive type=schema
//ff:what 사이트 단위 제출 자격 — IndexNow 키와 GSC SA 자격(인라인 JSON·파일 경로). 빈 필드는 전역 env fallback
//ff:why 멀티사이트 격리: 사이트 행 값 > 전역 env 우선순위를 시그니처로 표현 — 빈 Keys는 기존 단일 사이트 env 동작 그대로 (Phase020)
package archive

// Keys carries the per-site submission credentials. An empty field falls
// back to the instance-global environment variable of the same meaning, so
// Keys{} keeps the single-site behaviour: INDEXNOW_KEY for IndexNowKey,
// GSC_SA_JSON / GSC_SA_JSON_PATH for the service-account pair. A caller
// value always wins over the environment — a per-site path must never be
// shadowed by a global inline credential.
type Keys struct {
	IndexNowKey   string // IndexNow submission key; empty → INDEXNOW_KEY
	GSCSAJSON     string // inline service-account JSON; wins over GSCSAJSONPath
	GSCSAJSONPath string // service-account JSON file path; both empty → GSC_SA_JSON, then GSC_SA_JSON_PATH
}
