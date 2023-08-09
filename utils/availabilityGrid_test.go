package utils_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestGetAvailabilityNumbersFromCoordinates(t *testing.T) {
	o, e, v, p := constants.AvailabilityNumberOccupied, constants.AvailabilityNumberEmpty, constants.AvailabilityNumberValid, constants.AvailabilityNumberPlayable
	tests := []struct {
		name             string
		coordinates      []types.Coordinate
		availabilityGrid types.AvailabilityGrid
		want             []types.AvailabilityNumber
	}{
		{
			name: "Finds all four coordinates.",
			coordinates: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 0,
					Y: 2,
				},
				{
					X: 2,
					Y: 0,
				},
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{o, o, o},
				{p, o, o},
			},
			want: []types.AvailabilityNumber{
				o,
				e,
				v,
				p,
			},
		},
		{
			name: "Finds one coordinate.",
			coordinates: []types.Coordinate{
				{
					X: 1,
					Y: 1,
				},
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{o, p, o},
				{e, o, o},
			},
			want: []types.AvailabilityNumber{
				p,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetAvailabilityNumbersFromCoordinates(tt.coordinates, tt.availabilityGrid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestGetAvailabilityNumberForCoordinate(t *testing.T) {
	o, e, v, p := constants.AvailabilityNumberOccupied, constants.AvailabilityNumberEmpty, constants.AvailabilityNumberValid, constants.AvailabilityNumberPlayable
	tests := []struct {
		name             string
		coordinate       types.Coordinate
		availabilityGrid types.AvailabilityGrid
		want             types.AvailabilityNumber
	}{
		{
			name: "Finds occupied",
			coordinate: types.Coordinate{
				X: 0,
				Y: 0,
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{e, p, v},
				{v, v, v},
			},
			want: o,
		},
		{
			name: "Finds empty",
			coordinate: types.Coordinate{
				X: 0,
				Y: 1,
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{e, p, v},
				{v, v, v},
			},
			want: e,
		},
		{
			name: "Finds valid",
			coordinate: types.Coordinate{
				X: 2,
				Y: 1,
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{e, p, v},
				{v, v, v},
			},
			want: v,
		},
		{
			name: "Finds playable",
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			availabilityGrid: types.AvailabilityGrid{
				{o, e, v},
				{e, p, v},
				{v, v, v},
			},
			want: p,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetAvailabilityNumberForCoordinate(tt.coordinate, tt.availabilityGrid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestGetCoordiantesWithMinAvailabilityNumber(t *testing.T) {
	o, e, v, p := constants.AvailabilityNumberOccupied, constants.AvailabilityNumberEmpty, constants.AvailabilityNumberValid, constants.AvailabilityNumberPlayable
	tests := []struct {
		name               string
		availabilityNumber types.AvailabilityNumber
		availabilityGrid   types.AvailabilityGrid
		want               []types.Coordinate
	}{
		{
			name:               "Finds all playable",
			availabilityNumber: p,
			availabilityGrid: types.AvailabilityGrid{
				{p, e, o},
				{e, e, p},
				{v, v, v},
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
					Y: 2,
				},
			},
		},
		{
			name:               "Finds empty, valid, and playable.",
			availabilityNumber: e,
			availabilityGrid: types.AvailabilityGrid{
				{p, e, o},
				{o, e, p},
				{v, o, v},
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 1,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetCoordiantesWithMinAvailabilityNumber(tt.availabilityNumber, tt.availabilityGrid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestGetAvailabilityGridFromGrid(t *testing.T) {
	r, g, b, y, n := constants.ColorNumberRed, constants.ColorNumberGreen, constants.ColorNumberBlue, constants.ColorNumberYellow, constants.ColorNumberEmpty
	o, e, v, p := constants.AvailabilityNumberOccupied, constants.AvailabilityNumberEmpty, constants.AvailabilityNumberValid, constants.AvailabilityNumberPlayable
	tests := []struct {
		name        string
		colorNumber types.ColorNumber
		grid        types.Grid
		want        types.AvailabilityGrid
	}{
		{
			name:        "Finds for red",
			colorNumber: r,
			grid: types.Grid{
				{r, n, b},
				{g, n, n},
				{y, n, n},
			},
			want: types.AvailabilityGrid{
				{o, e, o},
				{o, p, v},
				{o, v, v},
			},
		},
		{
			name:        "All Occupied",
			colorNumber: r,
			grid: types.Grid{
				{r, r, b},
				{g, b, b},
				{y, g, y},
			},
			want: types.AvailabilityGrid{
				{o, o, o},
				{o, o, o},
				{o, o, o},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetAvailabilityGridFromGrid(tt.colorNumber, tt.grid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestCreateAvailabilityForCoordinate(t *testing.T) {
	r, g, b, y, n := constants.ColorNumberRed, constants.ColorNumberGreen, constants.ColorNumberBlue, constants.ColorNumberYellow, constants.ColorNumberEmpty
	o, e, v, p := constants.AvailabilityNumberOccupied, constants.AvailabilityNumberEmpty, constants.AvailabilityNumberValid, constants.AvailabilityNumberPlayable
	tests := []struct {
		name        string
		coordinate  types.Coordinate
		colorNumber types.ColorNumber
		grid        types.Grid
		want        types.AvailabilityNumber
	}{
		{
			name: "Finds playable",
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			colorNumber: r,
			grid: types.Grid{
				{g, r, b},
				{n, n, r},
				{y, n, n},
			},
			want: p,
		},
		{
			name: "Finds occupied",
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: r,
			grid: types.Grid{
				{n, n, n},
				{n, g, n},
				{n, n, n},
			},
			want: o,
		},
		{
			name: "Finds empy",
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: r,
			grid: types.Grid{
				{r, r, r},
				{r, n, r},
				{r, r, r},
			},
			want: e,
		},
		{
			name: "Finds valid",
			coordinate: types.Coordinate{
				X: 1,
				Y: 1,
			},
			colorNumber: r,
			grid: types.Grid{
				{b, g, b},
				{y, n, b},
				{b, b, b},
			},
			want: v,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.CreateAvailabilityForCoordinate(tt.coordinate, tt.colorNumber, tt.grid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestCreateEmptyAvailabilityGrid(t *testing.T) {
	e := constants.AvailabilityNumberEmpty
	tests := []struct {
		name   string
		width  int
		height int
		want   types.AvailabilityGrid
	}{
		{
			name:   "5x5",
			width:  5,
			height: 5,
			want: types.AvailabilityGrid{
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
			},
		},
		{
			name:   "1x5",
			width:  1,
			height: 5,
			want: types.AvailabilityGrid{
				{e, e, e, e, e},
			},
		},
		{
			name:   "3x1",
			width:  3,
			height: 1,
			want: types.AvailabilityGrid{
				{e},
				{e},
				{e},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.CreateEmptyAvailabilityGrid(tt.width, tt.height)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestGetAvailabilityGridSize(t *testing.T) {
	e := constants.AvailabilityNumberEmpty
	tests := []struct {
		name             string
		availabilityGrid types.AvailabilityGrid
		want             [2]int
	}{
		{
			name: "5x5",
			availabilityGrid: types.AvailabilityGrid{
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
				{e, e, e, e, e},
			},
			want: [2]int{5, 5},
		},
		{
			name: "2x6",
			availabilityGrid: types.AvailabilityGrid{
				{e, e, e, e, e, e},
				{e, e, e, e, e, e},
			},
			want: [2]int{2, 6},
		},
		{
			name: "3x2",
			availabilityGrid: types.AvailabilityGrid{
				{e, e},
				{e, e},
				{e, e},
			},
			want: [2]int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetAvailabilityGridSize(tt.availabilityGrid)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}
