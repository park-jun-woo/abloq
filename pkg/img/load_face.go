//ff:func feature=image type=parser control=sequence
//ff:what TTF/OTF 폰트를 지정 크기 Face로 로드 — 경로 미지정 시 임베디드 Go Bold (라틴 전용)
package img

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/opentype"
)

// LoadFace parses path (or the embedded Go Bold when path is empty) at size px.
func LoadFace(path string, size float64) (font.Face, error) {
	data := gobold.TTF
	if path != "" {
		var err error
		data, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
	}
	f, err := opentype.Parse(data)
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(f, &opentype.FaceOptions{Size: size, DPI: 72, Hinting: font.HintingFull})
}
