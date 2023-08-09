package utils_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestColorNumberInColorNumbers(t *testing.T) {
	r, g, b, y, e := constants.ColorNumberRed, constants.ColorNumberGreen, constants.ColorNumberBlue, constants.ColorNumberYellow, constants.ColorNumberEmpty
	tests := []struct {
		name         string
		colorNumber  types.ColorNumber
		colorNumbers []types.ColorNumber
		want         bool
	}{
		{
			name:        "Color numbers has red.",
			colorNumber: r,
			colorNumbers: []types.ColorNumber{
				b,
				y,
				g,
				e,
				r,
			},
			want: true,
		},
		{
			name:        "Color numbers doesn't have red.",
			colorNumber: r,
			colorNumbers: []types.ColorNumber{
				b,
				y,
				g,
				e,
			},
			want: false,
		},
		{
			name:         "Returns false when empty slice.",
			colorNumber:  r,
			colorNumbers: []types.ColorNumber{},
			want:         false,
		},
		{
			name:        "Returns true when with empty color.",
			colorNumber: e,
			colorNumbers: []types.ColorNumber{
				e,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ColorNumberInColorNumbers(tt.colorNumber, tt.colorNumbers)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}
