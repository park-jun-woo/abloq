//ff:func feature=cli type=command control=sequence
//ff:what archiveKeysFromEnv가 INDEXNOW_KEY·GSC_SA_JSON·GSC_SA_JSON_PATH env를 그대로 archive.Keys 필드에 옮기는지 검증
package main

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/archive"
)

func TestArchiveKeysFromEnv(t *testing.T) {
	t.Setenv("INDEXNOW_KEY", "k123")
	t.Setenv("GSC_SA_JSON", `{"client_email":"x@y"}`)
	t.Setenv("GSC_SA_JSON_PATH", "/secrets/sa.json")
	keys := archiveKeysFromEnv()
	if keys.IndexNowKey != "k123" || keys.GSCSAJSON != `{"client_email":"x@y"}` || keys.GSCSAJSONPath != "/secrets/sa.json" {
		t.Errorf("keys = %+v, want the three env values verbatim", keys)
	}

	t.Setenv("INDEXNOW_KEY", "")
	t.Setenv("GSC_SA_JSON", "")
	t.Setenv("GSC_SA_JSON_PATH", "")
	if keys = archiveKeysFromEnv(); keys != (archive.Keys{}) {
		t.Errorf("unset env must yield empty Keys (env fallback inside pkg/archive): %+v", keys)
	}
}
