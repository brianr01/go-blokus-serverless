package utils_test

import (
	"image/color"
	"testing"

	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestIsPixelWhite(t *testing.T) {
	w := color.RGBA{255, 255, 255, 0xff}
	b := color.RGBA{0, 0, 0, 0xff}
	tests := []struct {
		name   string
		x      int
		y      int
		colors [][]color.RGBA
		want   bool
	}{
		{
			name: "with white",
			x:    1,
			y:    0,
			colors: [][]color.RGBA{
				{b, b},
				{w, b},
			},
			want: true,
		},
		{
			name: "with black",
			x:    1,
			y:    0,
			colors: [][]color.RGBA{
				{w, w},
				{b, w},
			},
			want: false,
		},
		{
			name: "with white in corner",
			x:    1,
			y:    1,
			colors: [][]color.RGBA{
				{w, w},
				{b, w},
			},
			want: true,
		},
		{
			name: "with black in corner",
			x:    1,
			y:    1,
			colors: [][]color.RGBA{
				{w, w},
				{b, b},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := utils.CreateImageFrom2dColor(tt.colors)
			result := utils.IsPixelWhite(tt.x, tt.y, img)
			if !result == tt.want {
				t.Errorf(
					"Test '%s' did not retun the expected output.\n Want: %v\nResult: %v\n",
					tt.name,
					tt.want,
					result,
				)
			}
		})
	}
}
