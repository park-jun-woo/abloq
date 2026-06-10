//ff:func feature=cli type=output control=iteration dimension=1
//ff:what 기록된 파생물 경로를 한 줄씩 출력
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/abloq/pkg/gen"
)

// printOutputs lists each written derived file path on its own line.
func printOutputs(out io.Writer, dir string, outs []gen.Output) {
	for _, o := range outs {
		fmt.Fprintln(out, filepath.Join(dir, o.Path))
	}
}
