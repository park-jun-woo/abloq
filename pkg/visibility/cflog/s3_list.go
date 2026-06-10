//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what ListObjectsV2 페이지네이션 전체 순회 — 소스 Prefix+prefix 아래 키를 오름차순으로, afterKey는 start-after로 전달
package cflog

import (
	"encoding/xml"
	"net/url"
	"sort"
	"strings"
)

// List walks the bucket's ListObjectsV2 pages and returns every key under
// the configured Prefix plus the call's prefix, ascending, strictly after
// afterKey. The ingest cursor never rides in afterKey (the cursor is an
// hour boundary, not a key) — it exists for listing economy only.
func (s S3Source) List(prefix, afterKey string) ([]string, error) {
	var keys []string
	token := ""
	for {
		q := url.Values{}
		q.Set("list-type", "2")
		q.Set("prefix", s.Prefix+prefix)
		if afterKey != "" {
			q.Set("start-after", s.Prefix+afterKey)
		}
		if token != "" {
			q.Set("continuation-token", token)
		}
		resp, err := s.s3Do(s.baseURL() + "/?" + q.Encode())
		if err != nil {
			return nil, err
		}
		var page s3ListResult
		err = xml.NewDecoder(resp.Body).Decode(&page)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		for _, k := range page.Keys {
			keys = append(keys, strings.TrimPrefix(k, s.Prefix))
		}
		if !page.IsTruncated || page.NextContinuationToken == "" {
			break
		}
		token = page.NextContinuationToken
	}
	sort.Strings(keys)
	return keys, nil
}
