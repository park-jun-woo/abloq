//ff:func feature=quest type=parser control=iteration dimension=1 topic=queue
//ff:what Seed 공통 본체 — quests/queue/ 디렉토리 스캔, kind 일치 큐 파일 1개=Item 1개(Key=큐 파일 키), priority 내림차순 정렬 (소비 퀘스트 3종 공유)
//ff:why payload는 Seed 시점에 Item으로 고정한다(작업트리 큐 파일 재독 금지 — 변조 치즈). 우선순위 내림차순은 백엔드 스코어러가 매긴 운용 순서의 박제다 (Phase018 계획)
package common

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/park-jun-woo/reins/pkg/quest"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// SeedQueue scans the instance's quests/queue/ directory and seeds one TODO
// item per queue file whose kind matches, highest priority first (join key
// ascending on ties). The optional single argument is the instance directory
// (default "."); the nearest ancestor holding blog.yaml is the root. The
// queue payload is frozen into the item at seed time.
func SeedQueue(kind string, args []string) ([]*quest.Item, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("usage: scan [instance-dir]")
	}
	dir := "."
	if len(args) == 1 {
		dir = args[0]
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	root, err := FindRoot(abs)
	if err != nil {
		return nil, err
	}
	b, diags, err := blogyaml.Load(filepath.Join(root, "blog.yaml"))
	if err != nil {
		return nil, err
	}
	if len(diags) > 0 {
		return nil, fmt.Errorf("blog.yaml: %s", diags[0].String())
	}
	queueDir := filepath.Join(root, "quests", "queue")
	entries, err := os.ReadDir(queueDir)
	if err != nil {
		return nil, fmt.Errorf("queue directory unreadable (run the exporter or abloq scan first): %w", err)
	}
	var qits []queueio.Item
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".yaml") {
			continue
		}
		qit, err := readQueueFile(filepath.Join(queueDir, e.Name()))
		if err != nil {
			return nil, err
		}
		if qit.Kind != kind {
			continue
		}
		qits = append(qits, qit)
	}
	sort.SliceStable(qits, func(i, j int) bool { return lessQueueItem(qits[i], qits[j]) })
	var items []*quest.Item
	for _, qit := range qits {
		it, err := seedQueueItem(root, qit, b.Languages)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}
