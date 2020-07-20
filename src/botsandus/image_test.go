package botsandus

import (
	"image/color"
	"testing"
)

var white = color.RGBA{0xff, 0xff, 0xff, 0xff}
var black = color.RGBA{0x00, 0x00, 0x00, 0xff}

func TestCreateImage(t *testing.T) {
	img := CreateImageFromBinaryString("11011")
	for i := 0; i < 8; i++ {
		if img.At(i, 0) != white {
			t.Error("No leading white spaces")
		}
	}
	if img.At(8, 0) != white {
		t.Error("No correct information")
	}
	if img.At(10, 0) != black {
		t.Error("No correct information")
	}
	for i := 56; i < 255; i++ {
		if img.At(i, 0) != white {
			t.Error("No leading white spaces")
		}
	}
}
