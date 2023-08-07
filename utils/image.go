package utils

import (
	"image"
	"image/color"
)

var white = color.RGBA{255, 255, 255, 0xff}

func IsPixelWhite(x int, y int, img image.Image) bool {
	r, g, b, _ := img.At(x, y).RGBA()

	// Get the rgb components of white.
	wr, wg, wb, _ := white.RGBA()

	return r == wr && g == wg && b == wb
}

func CreateImageFrom2dColor(clrs [][]color.RGBA) image.Image {
	width := len(clrs)
	height := len(clrs[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := range clrs {
		for j, clr := range clrs[i] {
			img.Set(i, j, clr)
		}
	}

	return img
}
