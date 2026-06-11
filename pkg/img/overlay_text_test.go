//ff:func feature=image type=generator control=sequence
//ff:what OverlayText가 임의 크기 캔버스에 제목 잉크를 올리고, 제목·브랜드 폰트 로드 오류를 각각 전파하는지 검증
package img

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"golang.org/x/image/font/gofont/gobold"
)

func TestOverlayText(t *testing.T) {
	dst := image.NewRGBA(image.Rect(0, 0, 800, 400))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)
	if err := OverlayText(dst, OGSpec{Title: "Overlay on any canvas", Brand: "abloq"}); err != nil {
		t.Fatalf("OverlayText: %v", err)
	}
	if !hasInk(dst) {
		t.Error("no non-white pixel found — title was not composited")
	}
	if err := OverlayText(dst, OGSpec{Title: "x", FontPath: "/nonexistent.ttf"}); err == nil {
		t.Error("OverlayText with missing font expected error, got nil")
	}

	// brand face failure: both loads share one path, so only a path that
	// reads fine once and fails on the second read reaches the second error
	// branch — a FIFO serving one valid TTF then garbage does exactly that.
	fifo := filepath.Join(t.TempDir(), "font.fifo")
	if err := exec.Command("mkfifo", fifo).Run(); err != nil {
		t.Skipf("mkfifo unavailable: %v", err)
	}
	go func() {
		f, err := os.OpenFile(fifo, os.O_WRONLY, 0)
		if err != nil {
			return
		}
		f.Write(gobold.TTF)
		f.Close()
		// wait until the title-face reader is gone, so the garbage lands in
		// the brand-face read instead of the drained first pipe
		for {
			probe, err := os.OpenFile(fifo, os.O_WRONLY|syscall.O_NONBLOCK, 0)
			if err != nil {
				break // ENXIO: no reader attached
			}
			probe.Close()
			time.Sleep(time.Millisecond)
		}
		if f, err = os.OpenFile(fifo, os.O_WRONLY, 0); err == nil {
			f.Write([]byte("not a font"))
			f.Close()
		}
	}()
	if err := OverlayText(dst, OGSpec{Title: "x", FontPath: fifo}); err == nil {
		t.Error("OverlayText with failing brand face expected error, got nil")
	}
}
