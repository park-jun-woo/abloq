//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what SigV4 파생 서명 키 — AWS4+secret에서 날짜→리전→서비스→aws4_request 순 HMAC 체인
package cflog

// signingKey derives the SigV4 signing key for one date/region/service
// scope.
func signingKey(secret, dateStamp, region, service string) []byte {
	kDate := hmacSHA256([]byte("AWS4"+secret), dateStamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	return hmacSHA256(kService, "aws4_request")
}
