//ff:func feature=cli type=command control=sequence
//ff:what insight match 실행 본체 — 명세 로드(에러 진단 시 중단), 본문 대조, 미출현 목록 출력, 섹션 불일치는 에러
//ff:why 미출현 claim은 에러가 아니다 — 출현 = 대응 보장이 아니므로 결과는 REVIEW 보조 자료로만 제시한다 (Phase015)
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/insight"
)

// runInsightMatch screens insightPath claims against articlePath's body.
// Schema errors and section mismatch return errors; missing claims do not.
func runInsightMatch(out io.Writer, insightPath, articlePath string) error {
	ins, errs, warns, err := insight.Load(insightPath)
	if err != nil {
		return err
	}
	printDiagsText(out, warns)
	if len(errs) > 0 {
		printDiagsText(out, errs)
		return fmt.Errorf("%s: %d issue(s) found", insightPath, len(errs))
	}
	article, err := os.ReadFile(articlePath)
	if err != nil {
		return err
	}
	if want := insight.PathFor(articlePath); filepath.Clean(insightPath) != filepath.Clean(want) {
		fmt.Fprintf(out, "note: insight file is not at the conventional path %s\n", want)
	}
	res := insight.Match(ins, articlePath, article)
	printInsightMatch(out, res, len(ins.Claims))
	if res.Section != ins.Section {
		return fmt.Errorf("section mismatch: insight declares %q but the article lives in %q", ins.Section, res.Section)
	}
	return nil
}
