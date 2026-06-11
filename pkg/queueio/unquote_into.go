//ff:func feature=queueio type=parser control=sequence
//ff:what 최상위 스칼라 라인의 strconv.Quote 값을 unquote해 대상 필드에 적재 (Deserialize 전용)
package queueio

import (
	"fmt"
	"strconv"
	"strings"
)

// unquoteInto strips the field prefix from one top-level serialized line and
// unquotes the remaining scalar into dst.
func unquoteInto(dst *string, ln, prefix string) error {
	v, err := strconv.Unquote(strings.TrimPrefix(ln, prefix))
	if err != nil {
		return fmt.Errorf("queue file: %s%w", prefix, err)
	}
	*dst = v
	return nil
}
