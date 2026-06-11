//ff:func feature=quest type=rule control=sequence
//ff:what citation-exists 어댑터 스킵 경로 검증 — Target.Offline=true면 깨진 인용이 있어도 전체 스킵(발동 없음)
package writing

import (
	"strings"
	"testing"
)

func TestAdaptCitationOffline(t *testing.T) {
	root := writeInstance(t)
	art, _ := passFixtures()
	cited := strings.Replace(art,
		"This body mentions the alpha anchor.",
		"This body mentions the alpha anchor ([Broken Ref](http://127.0.0.1:1/gone)).",
		1)
	sub := subWith(t, root, cited, "")
	sub.Target.Offline = true
	if fired, _ := fireRule(t, adaptRule("citation-exists"), sub); fired {
		t.Error("citation-exists: fired despite Offline=true (skip path)")
	}
}
