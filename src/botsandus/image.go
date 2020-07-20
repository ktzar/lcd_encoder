package botsandus

import (
	"image"
	"image/color"
)

var imagePalette = []color.Color{
	color.RGBA{0xff, 0xff, 0xff, 0xff},
	color.RGBA{0x00, 0x00, 0x00, 0xff},
}

func CreateImageFromBinaryString(input string) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, 256, 1), imagePalette)
	for i, v := range input {
		if string(v) == "0" {
			img.Set(i+8, 0, imagePalette[1])
		}
	}
	return img
}
