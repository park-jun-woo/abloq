//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what GetObject 1회 — 소스 Prefix를 붙인 키의 본문 스트림 반환
package cflog

import "io"

// Get streams one object body (the caller closes it). Keys are relative to
// the configured Prefix, mirroring List.
func (s S3Source) Get(key string) (io.ReadCloser, error) {
	resp, err := s.s3Do(s.baseURL() + "/" + s3EscapePath(s.Prefix+key))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
