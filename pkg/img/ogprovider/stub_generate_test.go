//ff:func feature=image type=client control=sequence
//ff:what Stub.Generate가 고정 크기 캔버스(기본 1024×1024)와 설정된 에러를 반환하는지 검증
package ogprovider

import (
	"context"
	"errors"
	"testing"
)

func TestStubGenerate(t *testing.T) {
	m, err := Stub{W: 300, H: 200}.Generate(context.Background(), "p")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 300 || b.Dy() != 200 {
		t.Errorf("bounds = %v, want 300x200", b)
	}
	m, err = Stub{}.Generate(context.Background(), "p")
	if err != nil || m.Bounds().Dx() != 1024 {
		t.Errorf("zero-size stub: bounds %v err %v, want 1024 default", m.Bounds(), err)
	}
	if _, err := (Stub{Err: errors.New("boom")}).Generate(context.Background(), "p"); err == nil {
		t.Error("Err stub expected error, got nil")
	}
}
