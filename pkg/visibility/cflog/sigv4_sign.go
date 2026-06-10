//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what AWS SigV4 GET 서명 — X-Amz-Date(+세션 토큰) 설정, 정준 요청→서명 문자열→파생 키 HMAC으로 Authorization 헤더 부착
//ff:why stdlib만으로 ListObjectsV2·GetObject 2개 호출을 서명한다(신규 의존성 0) — 정확성은 AWS 공개 테스트 벡터(iam ListUsers 예제) 단위 테스트로 검증 (Phase012)
package cflog

import (
	"net/http"
	"strings"
	"time"
)

// emptyPayloadHash is SHA-256 of the empty string — every request we sign
// is a bodyless GET.
const emptyPayloadHash = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

// signRequest signs req (a bodyless GET) with AWS Signature Version 4. It
// sets X-Amz-Date (and X-Amz-Security-Token when a session token is
// present), canonicalizes the request over every header already set, and
// attaches the Authorization header. payloadHash is the hex SHA-256 of the
// request body; callers that must send x-amz-content-sha256 (S3) set that
// header before calling.
func signRequest(req *http.Request, accessKey, secretKey, sessionToken, region, service string, now time.Time, payloadHash string) {
	amzDate := now.UTC().Format("20060102T150405Z")
	dateStamp := now.UTC().Format("20060102")
	req.Header.Set("X-Amz-Date", amzDate)
	if sessionToken != "" {
		req.Header.Set("X-Amz-Security-Token", sessionToken)
	}
	uri := req.URL.EscapedPath()
	if uri == "" {
		uri = "/"
	}
	headers, signedHeaders := canonicalHeaders(req)
	canonical := strings.Join([]string{
		req.Method, uri, canonicalQuery(req.URL.RawQuery), headers, signedHeaders, payloadHash,
	}, "\n")
	scope := dateStamp + "/" + region + "/" + service + "/aws4_request"
	stringToSign := strings.Join([]string{
		"AWS4-HMAC-SHA256", amzDate, scope, sha256Hex([]byte(canonical)),
	}, "\n")
	signature := sha256HexMAC(signingKey(secretKey, dateStamp, region, service), stringToSign)
	req.Header.Set("Authorization",
		"AWS4-HMAC-SHA256 Credential="+accessKey+"/"+scope+
			", SignedHeaders="+signedHeaders+", Signature="+signature)
}
