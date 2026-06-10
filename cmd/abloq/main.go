//ff:func feature=cli type=command control=sequence
//ff:what abloq CLI 진입점 — 루트 명령 실행, 실패 시 exit 1
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "abloq:", err)
		os.Exit(1)
	}
}
