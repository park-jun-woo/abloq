//ff:func feature=cli type=command control=sequence topic=crawl
//ff:what crawl 수집 실행 본체 — 소스의 전 로그 키를 커서·마진 없이 집계해 히트·미지 봇·원시 카운터 출력
//ff:why CLI는 단발 분석이라 증분 상태가 없다 — 시간 프리픽스가 있는 키 전부를 통째로 읽는다. 원시 봇 카운터는 analyze-stats.py 대조 지점(이식 정확성 검증의 비교 단계) (Phase012)
package main

import (
	"io"

	"github.com/park-jun-woo/abloq/pkg/visibility/cflog"
)

// runIngestCrawl aggregates every log object of the source in one stateless
// pass: the article hit rows, the unknown-bot candidates and the unfiltered
// per-bot raw counters land on out.
func runIngestCrawl(out io.Writer, source, repo string) error {
	b, err := loadValidBlog(out, repo)
	if err != nil {
		return err
	}
	urls, err := cflog.BuildURLMap(repo, b)
	if err != nil {
		return err
	}
	src, err := cflog.OpenSource(source)
	if err != nil {
		return err
	}
	keys, err := src.List("", "")
	if err != nil {
		return err
	}
	logKeys := cflog.LogKeys(keys)
	agg, err := cflog.IngestKeys(src, urls, logKeys)
	if err != nil {
		return err
	}
	printCrawlAgg(out, len(logKeys), agg)
	return nil
}
