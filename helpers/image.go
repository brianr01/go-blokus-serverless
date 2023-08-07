package helpers

import (
	"image"
	"image/color"
)

var white = color.RGBA{255, 255, 255, 0xff}

func IsPixelWhite(x int, y int, image image.Image) bool {
	r, g, b, _ := image.At(x, y).RGBA()

	// Get the rgb components of white.
	wr, wg, wb, _ := white.RGBA()

	return r == wr && g == wg && b == wb
}
