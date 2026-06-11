//ff:func feature=image type=generator control=sequence
//ff:what ogGenerateOne 검증 — stub Provider 결과의 1200×630 정규화 WebP 기록, 오버레이 유무 바이트 차이, Provider 실패 전파
package img

import (
	"context"
	"errors"
	"image/color"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/image/webp"
)

func TestOGGenerateOne(t *testing.T) {
	dir := t.TempDir()
	blue := stubOGProvider{w: 900, h: 900, c: color.NRGBA{10, 10, 200, 255}}
	spec := OGAISpec{Slug: "post", Title: "T", Brand: "B"}

	// plain: provider image is cropped/flattened to 1200x630 and written
	plainOut := filepath.Join(dir, "plain.webp")
	job := ogJob{variant: OGVariant{Name: "a", Prompt: "p", Provider: blue}, n: 1, out: plainOut}
	if err := ogGenerateOne(context.Background(), spec, job); err != nil {
		t.Fatalf("plain: %v", err)
	}
	f, err := os.Open(plainOut)
	if err != nil {
		t.Fatalf("output missing: %v", err)
	}
	m, err := webp.Decode(f)
	f.Close()
	if err != nil {
		t.Fatalf("webp decode: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 1200 || b.Dy() != 630 {
		t.Errorf("bounds = %v, want 1200x630", b)
	}

	// overlay: title/brand composition changes the bytes
	overlayOut := filepath.Join(dir, "overlay.webp")
	job = ogJob{variant: OGVariant{Name: "a", Prompt: "p", Overlay: true, Provider: blue}, n: 1, out: overlayOut}
	if err := ogGenerateOne(context.Background(), spec, job); err != nil {
		t.Fatalf("overlay: %v", err)
	}
	plain, _ := os.ReadFile(plainOut)
	overlaid, _ := os.ReadFile(overlayOut)
	if string(plain) == string(overlaid) {
		t.Error("overlay output equals plain output — text was not composited")
	}

	// overlay failure (missing font) propagates
	badFont := OGAISpec{Slug: "post", Title: "T", FontPath: "/nonexistent.ttf"}
	job = ogJob{variant: OGVariant{Name: "a", Prompt: "p", Overlay: true, Provider: blue}, n: 1, out: filepath.Join(dir, "font.webp")}
	if err := ogGenerateOne(context.Background(), badFont, job); err == nil {
		t.Error("missing overlay font must error")
	}

	// provider failure propagates, nothing is written
	boom := stubOGProvider{err: errors.New("quota")}
	failOut := filepath.Join(dir, "fail.webp")
	job = ogJob{variant: OGVariant{Name: "bad", Prompt: "p", Provider: boom}, n: 1, out: failOut}
	if err := ogGenerateOne(context.Background(), spec, job); err == nil || !errors.Is(err, boom.err) {
		t.Errorf("failure: err = %v, want quota", err)
	}
	if _, err := os.Stat(failOut); err == nil {
		t.Error("failed job must not leave an output file")
	}
}
