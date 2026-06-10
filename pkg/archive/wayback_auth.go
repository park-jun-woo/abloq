//ff:func feature=archive type=client control=sequence
//ff:what SPN2 인증 헤더 조립 — WAYBACK_ACCESS_KEY/WAYBACK_SECRET_KEY 둘 다 있을 때만 "LOW key:secret"
package archive

import "os"

// waybackAuth builds the SPN2 "LOW <access>:<secret>" Authorization value.
// Without keys it returns "" — SPN2 then rejects the call and the rejection
// is recorded honestly as a failed receipt.
func waybackAuth() string {
	access := os.Getenv("WAYBACK_ACCESS_KEY")
	secret := os.Getenv("WAYBACK_SECRET_KEY")
	if access == "" || secret == "" {
		return ""
	}
	return "LOW " + access + ":" + secret
}
