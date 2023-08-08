package utils_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestColorLocation(t *testing.T) {
	tests := []struct {
		name       string
		color      types.ColorNumber
		coordinate types.Coordinate
		grid       types.Grid
		want       types.Grid
	}{
		{
			name:  "base case",
			color: 1,
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			grid: types.Grid{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: types.Grid{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ColorLocation(tt.color, tt.coordinate, tt.grid)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("The result %v does not match %v", result, tt.want)
			}
		})
	}
}

// func TestGetPlayableCoordinatesForColor(t *testing.T) {
// 	e, r, g, b, y := constants.ColorNumberEmpty, constants.ColorNumberRed, constants.ColorNumberGreen, constants.ColorNumberBlue, constants.ColorNumberYellow
// 	tests := []struct {
// 		name        string
// 		grid        types.Grid
// 		colorNumber types.ColorNumber
// 		want        []types.Coordinate
// 	}{
// 		{
// 			name: "Returns no playable coordintates on empty grid.",
// 			grid: types.Grid{
// 				{e, e, e},
// 				{e, e, e},
// 				{e, e, e},
// 			},
// 			colorNumber: b,
// 			want:        []types.Coordinate{},
// 		},
// 		{
// 			name: "Returns no playable coordintates on full grid.",
// 			grid: types.Grid{
// 				{b, y, r},
// 				{g, y, b},
// 				{r, r, g},
// 			},
// 			colorNumber: b,
// 			want:        []types.Coordinate{},
// 		},
// 		{
// 			name: "Returns one playable coordintate",
// 			grid: types.Grid{
// 				{b, y, r},
// 				{g, e, r},
// 				{r, r, g},
// 			},
// 			colorNumber: b,
// 			want: []types.Coordinate{
// 				{
// 					X: 1,
// 					Y: 1,
// 				},
// 			},
// 		},
// 		{
// 			name: "Returns no playable coordintates with same color edge",
// 			grid: types.Grid{
// 				{b, y, r},
// 				{g, e, b},
// 				{r, r, g},
// 			},
// 			colorNumber: b,
// 			want: []types.Coordinate{},
// 		},
// 		{
// 			name: "Returns 4 playable coordintates",
// 			grid: types.Grid{
// 				{e, e, e},
// 				{e, b, e},
// 				{e, e, e},
// 			},
// 			colorNumber: b,
// 			want: []types.Coordinate{
// 				{
// 					X: 0,
// 					Y: 0,
// 				},
// 				{
// 					X: 0,
// 					Y: 2,
// 				},
// 				{
// 					X: 2,
// 					Y: 0,
// 				},
// 				{
// 					X: 2,
// 					Y: 2,
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := utils.GetPlayableCoordinatesForColor(tt.colorNumber, tt.grid)
// 			if !reflect.DeepEqual(tt.want, result) {
// 				t.Errorf("The result %v does not match %v", result, tt.want)
// 			}
// 		})
// 	}
// }

func TestIsPlayableCoordinate(t *testing.T) {
	e, r, g, b, y := constants.ColorNumberEmpty, constants.ColorNumberRed, constants.ColorNumberGreen, constants.ColorNumberBlue, constants.ColorNumberYellow
	tests := []struct {
		name        string
		grid        types.Grid
		coordinate  types.Coordinate
		colorNumber types.ColorNumber
		want        bool
	}{
		{
			name: "Not playable on empty board.",
			grid: types.Grid{
				{e, e, e},
				{e, e, e},
				{e, e, e},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: constants.ColorNumberBlue,
			want:        false,
		},
		{
			name: "No playable on full board.",
			grid: types.Grid{
				{b, g, r},
				{y, b, g},
				{y, b, r},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: constants.ColorNumberBlue,
			want:        false,
		},
		{
			name: "Square is playable with one valid corner.",
			grid: types.Grid{
				{e, e, e},
				{e, e, e},
				{b, e, e},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: b,
			want:        true,
		},
		{
			name: "Square is not playable with one invalid edge",
			grid: types.Grid{
				{e, e, e},
				{e, e, e},
				{b, b, e},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: b,
			want:        false,
		},
		{
			name: "Square is playable with one valid edges",
			grid: types.Grid{
				{e, y, e},
				{g, e, y},
				{b, r, e},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: b,
			want:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsPlayableCoordinate(tt.coordinate, tt.colorNumber, tt.grid)
			if tt.want != result {
				t.Errorf("The result %v does not match %v", result, tt.want)
			}
		})
	}
}

func TestGetEdgeCoordinatesFromCoordinate(t *testing.T) {
	grid4x4 := types.Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	tests := []struct {
		name       string
		grid       types.Grid
		coordinate types.Coordinate
		want       []types.Coordinate
	}{
		{
			name: "Returns 4 edges",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 1,
				},
				{
					X: 1,
					Y: 2,
				},
			},
		},
		{
			name: "Returns +x and +y edges",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 0,
				Y: 0,
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
			},
		},
		{
			name: "Returns -x and -y edges",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 3,
				Y: 3,
			},
			want: []types.Coordinate{
				{
					X: 2,
					Y: 3,
				},
				{
					X: 3,
					Y: 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetEdgeCoordinatesFromCoordinate(tt.coordinate, tt.grid)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("The expected and result do not match.\n Expected: %v\n Actual:   %v\n", tt.want, result)
			}
		})
	}
}

func TestGetCornerCoordinatesFromCoordinate(t *testing.T) {
	grid4x4 := types.Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	tests := []struct {
		name       string
		grid       types.Grid
		coordinate types.Coordinate
		want       []types.Coordinate
	}{
		{
			name: "Returns 4 corners",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: 2,
				},
				{
					X: 2,
					Y: 0,
				},
				{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			name: "Returns +x +y corner",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 0,
				Y: 0,
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 1,
				},
			},
		},
		{
			name: "Returns -x -y corner",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 3,
				Y: 3,
			},
			want: []types.Coordinate{
				{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			name: "Returns +x -y corner",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 0,
				Y: 3,
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 2,
				},
			},
		},
		{
			name: "Returns -x +y corner",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 3,
				Y: 0,
			},
			want: []types.Coordinate{
				{
					X: 2,
					Y: 1,
				},
			},
		},
		{
			name: "Returns (+x +y) and (+x, -y) corner",
			grid: grid4x4,
			coordinate: types.Coordinate{
				X: 0,
				Y: 1,
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 1,
					Y: 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetCornerCoordinatesFromCoordinate(tt.coordinate, tt.grid)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("The expected and result do not match.\n Expected: %v\n Actual:   %v\n", tt.want, result)
			}
		})
	}
}

func TestIsCoordinateOnGrid(t *testing.T) {
	grid3x3 := types.Grid{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	tests := []struct {
		name       string
		grid       types.Grid
		coordinate types.Coordinate
		want       bool
	}{
		{
			name: "Coordinate in middle of the grid.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			want: true,
		},
		{
			name: "Coordinate on edge of grid in +x direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 2,
				Y: 1,
			},
			want: true,
		},
		{
			name: "Coordinate on edge of grid in -x direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 0,
				Y: 1,
			},
			want: true,
		},
		{
			name: "Coordinate on edge of grid in +y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 1,
				Y: 2,
			},
			want: true,
		},
		{
			name: "Coordinate on edge of grid in -y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			want: true,
		},
		{
			name: "Coordinate on corner of grid in +x and +y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 2,
				Y: 2,
			},
			want: true,
		},
		{
			name: "Coordinate on corner of grid in -x and -y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 0,
				Y: 0,
			},
			want: true,
		},
		{
			name: "Coordinate on corner of grid in +x and -y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 2,
				Y: 0,
			},
			want: true,
		},
		{
			name: "Coordinate on corner of grid in -x and +y direction",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 0,
				Y: 2,
			},
			want: true,
		},
		{
			name: "Coordinate off grid in +x direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 3,
				Y: 0,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in -x direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: -1,
				Y: 0,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in +y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 0,
				Y: 3,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in -y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 0,
				Y: -1,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in +x and +y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 3,
				Y: 3,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in -x and -y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: -1,
				Y: -1,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in +x and -y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: 3,
				Y: -1,
			},
			want: false,
		},
		{
			name: "Coordinate off grid in -x and +y direction.",
			grid: grid3x3,
			coordinate: types.Coordinate{
				X: -1,
				Y: 3,
			},
			want: false,
		},
		{
			name: "Coordinate on recangular grid.",
			grid: types.Grid{
				{1, 1},
			},
			coordinate: types.Coordinate{
				X: 0,
				Y: 1,
			},
			want: true,
		},
		{
			name: "Coordinate off recangular grid.",
			grid: types.Grid{
				{1, 1},
			},
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsCoordinateOnGrid(tt.coordinate, tt.grid)
			if tt.want != result {
				t.Errorf("The result %v does not match expected %v.", result, tt.want)
			}
		})
	}
}

func TestGetGridSize(t *testing.T) {
	tests := []struct {
		name string
		grid types.Grid
		want [2]int
	}{
		{
			name: "1x1 Grid",
			grid: types.Grid{
				{0},
			},
			want: [2]int{1, 1},
		},
		{
			name: "10x20 Grid",
			grid: types.Grid{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: [2]int{10, 20},
		},
		{
			name: "9x4 Grid",
			grid: types.Grid{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			want: [2]int{9, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetGridSize(tt.grid)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("The grid size %v did not match the expected grid size of %v", result, tt.want)
			}
		})
	}
}
