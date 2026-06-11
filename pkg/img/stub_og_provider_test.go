//ff:type feature=image type=client
//ff:what 테스트 stub Provider — 고정 크기·단색 캔버스 또는 지정 에러 (img 내부용, ogprovider 순환 import 회피)
package img

import (
	"image/color"
)

type stubOGProvider struct {
	w, h int
	c    color.Color
	err  error
}
