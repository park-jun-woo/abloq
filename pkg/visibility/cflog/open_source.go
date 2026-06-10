//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 소스 명세 해석 — s3://bucket/prefix면 env 자격증명의 S3 소스, 그 외는 로컬 디렉토리 소스
//ff:why 자격증명은 전부 백엔드 env에만(설계서 §3.3) — 테스트·Hurl·CLI는 디렉토리 소스만 쓰고 실 S3는 본번 판정(Phase019 이월) (Phase012)
package cflog

import (
	"fmt"
	"os"
	"strings"
)

// OpenSource resolves a log source spec: "s3://bucket/prefix" builds an
// S3Source from the standard AWS env credentials, anything else must be an
// existing local directory.
func OpenSource(spec string) (Source, error) {
	if rest, isS3 := strings.CutPrefix(spec, "s3://"); isS3 {
		return openS3Source(spec, rest)
	}
	fi, err := os.Stat(spec)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("cflog: source %q is not a directory", spec)
	}
	return DirSource{Root: spec}, nil
}
