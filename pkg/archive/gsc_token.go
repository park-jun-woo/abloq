//ff:func feature=archive type=client control=sequence
//ff:what 빈 Keys로 GSCTokenWith에 위임 — GSC_SA_JSON/GSC_SA_JSON_PATH 전역 env 자격을 쓰는 단일 사이트 하위호환 진입점
package archive

// Google OAuth2 scopes the abloq backends ask for: the archiver publishes
// through the Indexing API, the visibility poller reads Search Console.
const (
	ScopeIndexing           = "https://www.googleapis.com/auth/indexing"
	ScopeWebmastersReadonly = "https://www.googleapis.com/auth/webmasters.readonly"
)

// GSCToken keeps the single-site entrypoint: empty Keys load the
// service-account credentials from GSC_SA_JSON / GSC_SA_JSON_PATH.
func GSCToken(scope string) (string, error) {
	return GSCTokenWith(Keys{}, scope)
}
