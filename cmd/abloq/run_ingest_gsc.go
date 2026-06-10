//ff:func feature=cli type=command control=sequence topic=gsc
//ff:what gsc 수집 실행 본체 — SA 토큰(webmasters.readonly)으로 최근 N일 닫힌 일자의 Search Analytics를 조회해 출력 (DB 없이)
//ff:why CLI는 단발 분석이라 커서가 없다 — 빈 커서 + lookback=days로 같은 pkg를 무상태 공유한다. 자격은 env로만(GSC_SA_JSON·GSC_TOKEN_URL·GSC_SEARCH_API_BASE) (Phase013)
package main

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/park-jun-woo/abloq/pkg/archive"
	"github.com/park-jun-woo/abloq/pkg/visibility/gsc"
)

// runIngestGSC fetches the last days closed days of Search Analytics rows
// for the property and prints them — API output only, no database. An empty
// site falls back to the blog.yaml baseURL of repo (a URL-prefix property).
func runIngestGSC(out io.Writer, site, repo string, days int) error {
	if site == "" {
		b, err := loadValidBlog(out, repo)
		if err != nil {
			return err
		}
		site = b.Site.BaseURL
		if !strings.HasSuffix(site, "/") {
			site += "/"
		}
	}
	token, err := archive.GSCToken(archive.ScopeWebmastersReadonly)
	if err != nil {
		return err
	}
	base := os.Getenv("GSC_SEARCH_API_BASE")
	if base == "" {
		base = "https://searchconsole.googleapis.com"
	}
	dates := gsc.Dates("", time.Now().UTC(), 2, days)
	res, err := gsc.Collect(base, token, site, dates)
	if err != nil {
		return err
	}
	printGscRows(out, res)
	return nil
}
