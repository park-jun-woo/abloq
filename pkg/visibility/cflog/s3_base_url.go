//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what S3 소스의 베이스 URL — Endpoint 오버라이드가 없으면 가상 호스트 형식 https://{bucket}.s3.{region}.amazonaws.com
package cflog

import "fmt"

// baseURL resolves the bucket endpoint: the test stub override wins,
// otherwise the regional virtual-hosted-style URL.
func (s S3Source) baseURL() string {
	if s.Endpoint != "" {
		return s.Endpoint
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com", s.Bucket, s.Region)
}
