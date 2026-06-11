//ff:type feature=image type=client
//ff:what 테스트 stub provider — 고정 크기 단색 캔버스 반환(네트워크 0), Err 지정 시 실패 재현
package ogprovider

import "image/color"

// Stub is the test provider: Generate returns a W×H solid canvas without any
// networking, or Err when set (partial-failure scenarios). Zero-size stubs
// default to 1024×1024.
type Stub struct {
	W, H  int
	Color color.NRGBA
	Err   error
}
