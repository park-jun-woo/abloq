//ff:type feature=archive type=schema
//ff:what GSC 서비스 계정 자격 — client_email과 RSA private_key (Google SA JSON의 부분집합)
package archive

// serviceAccount is the subset of a Google service-account JSON the
// Indexing API token exchange needs.
type serviceAccount struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}
