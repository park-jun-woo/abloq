//ff:func feature=quest type=parser control=sequence topic=queue
//ff:what 큐 파일 1개 읽기+Deserialize — 실패는 파일 경로를 병기한 에러 (SeedQueue 전용)
package common

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/abloq/pkg/queueio"
)

// readQueueFile loads one serialized queue file from disk.
func readQueueFile(path string) (queueio.Item, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return queueio.Item{}, err
	}
	qit, err := queueio.Deserialize(data)
	if err != nil {
		return queueio.Item{}, fmt.Errorf("%s: %w", path, err)
	}
	return qit, nil
}
