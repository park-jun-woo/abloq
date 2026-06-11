//ff:type feature=image type=client
//ff:what OG 배경 이미지 provider 인터페이스 — 프롬프트 1개로 임의 크기 이미지 생성 (stdlib 타입만, 구현은 ogprovider)
//ff:why pkg/img은 네트워크 0 불변 — 인터페이스만 여기 두고 HTTP 구현은 pkg/img/ogprovider에 단방향 격리, local은 비경유(RenderOG 직행) (BUG002)
package img

import (
	"context"
	"image"
)

// Provider generates one raw background image (any size) for a prompt; the
// caller post-processes (center crop, flatten, overlay, WebP). Implementations
// live in pkg/img/ogprovider — img itself never imports them, keeping this
// package free of networking. The local deterministic card does not go
// through Provider at all.
type Provider interface {
	Generate(ctx context.Context, prompt string) (image.Image, error)
}
