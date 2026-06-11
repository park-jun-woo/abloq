//ff:func feature=cli type=command control=sequence
//ff:what 단일 사이트 CLI의 제출 자격 — INDEXNOW_KEY·GSC_SA_JSON·GSC_SA_JSON_PATH env를 archive.Keys로 묶어 명시 전달
//ff:why CLI는 사이트 행이 없다 — 기존 env 이름 그대로를 인자로 끌어올려 pkg/archive의 자격 주입 시그니처와 하위호환을 함께 지킨다 (Phase020)
package main

import (
	"os"

	"github.com/park-jun-woo/abloq/pkg/archive"
)

// archiveKeysFromEnv collects the single-site submission credentials from
// the environment the CLI has always used. Passing them explicitly keeps the
// CLI behaviour byte-identical while pkg/archive now takes keys as input.
func archiveKeysFromEnv() archive.Keys {
	return archive.Keys{
		IndexNowKey:   os.Getenv("INDEXNOW_KEY"),
		GSCSAJSON:     os.Getenv("GSC_SA_JSON"),
		GSCSAJSONPath: os.Getenv("GSC_SA_JSON_PATH"),
	}
}
