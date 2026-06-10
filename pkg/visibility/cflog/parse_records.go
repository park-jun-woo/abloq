//ff:func feature=visibility type=parser control=iteration dimension=1 topic=crawl
//ff:what gzip 로그 스트림 전체를 행 단위로 파싱해 Record 목록으로 — 탈락 행은 건너뜀
package cflog

import (
	"bufio"
	"compress/gzip"
	"io"
)

// parseRecords decompresses one .gz log object and parses every line,
// skipping the lines parseLine drops.
func parseRecords(r io.Reader) ([]Record, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var recs []Record
	sc := bufio.NewScanner(zr)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		if rec, ok := parseLine(sc.Text()); ok {
			recs = append(recs, rec)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return recs, nil
}
