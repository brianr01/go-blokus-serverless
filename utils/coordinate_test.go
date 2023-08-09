package utils_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestAddCoordinates(t *testing.T) {
	tests := []struct {
		name        string
		coordinates []types.Coordinate
		want        types.Coordinate
	}{
		{
			name: "2 coordinates add",
			coordinates: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
			},
			want: types.Coordinate{
				X: 1,
				Y: 1,
			},
		},
		{
			name: "2 coordinates with one negative",
			coordinates: []types.Coordinate{
				{
					X: -1,
					Y: 0,
				},
				{
					X: 2,
					Y: 1,
				},
			},
			want: types.Coordinate{
				X: 1,
				Y: 1,
			},
		},
		{
			name: "2 coordinates with both negative",
			coordinates: []types.Coordinate{
				{
					X: -1,
					Y: -2,
				},
				{
					X: -5,
					Y: -1,
				},
			},
			want: types.Coordinate{
				X: -6,
				Y: -3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.AddCoordinates(tt.coordinates...)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestFilterCoordinatesByMatchingCoordinates(t *testing.T) {
	tests := []struct {
		name                string
		coordinatesToFilter []types.Coordinate
		coordinatesToMatch  []types.Coordinate
		want                []types.Coordinate
	}{
		{
			name: "One coordinate matches.",
			coordinatesToFilter: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
			},
			coordinatesToMatch: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
			},
		},
		{
			name: "Mupltiple coordinates match.",
			coordinatesToFilter: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
				{
					X: 2,
					Y: 2,
				},
			},
			coordinatesToMatch: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 2,
				},
			},
			want: []types.Coordinate{
				{
					X: 1,
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
			result := utils.FilterCoordinatesByMatchingCoordinates(tt.coordinatesToFilter, tt.coordinatesToMatch)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestIsCoordinateInCoordinates(t *testing.T) {
	tests := []struct {
		name               string
		coordinate         types.Coordinate
		coordinatesToMatch []types.Coordinate
		want               bool
	}{
		{
			name: "Coordinate matches.",
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinatesToMatch: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
			},
			want: true,
		},
		{
			name: "Coordinate matches last.",
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinatesToMatch: []types.Coordinate{
				{
					X: 2,
					Y: 0,
				},
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
					Y: 0,
				},
			},
			want: true,
		},
		{
			name: "Coordinate doesn't matche.",
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinatesToMatch: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 2,
					Y: 0,
				},
				{
					X: 0,
					Y: 1,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsCoordinateInCoordinates(tt.coordinate, tt.coordinatesToMatch)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestShiftCoordinatesByCoordinate(t *testing.T) {
	tests := []struct {
		name        string
		coordinate  types.Coordinate
		coordinates []types.Coordinate
		want        []types.Coordinate
	}{
		{
			name: "Coordinates shift.",
			coordinate: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinates: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
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
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
				{
					X: 1,
					Y: 1,
				},
				{
					X: 2,
					Y: 1,
				},
			},
		},
		{
			name: "Coordinates shift with negative numbers.",
			coordinate: types.Coordinate{
				X: -1,
				Y: -2,
			},
			coordinates: []types.Coordinate{
				{
					X: 1,
					Y: 2,
				},
				{
					X: 1,
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
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 0,
					Y: -2,
				},
				{
					X: -1,
					Y: -1,
				},
				{
					X: 0,
					Y: -1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ShiftCoordinatesByCoordinate(tt.coordinate, tt.coordinates)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestAreCoordinatesEqual(t *testing.T) {
	tests := []struct {
		name        string
		coordinate1 types.Coordinate
		coordinate2 types.Coordinate
		want        bool
	}{
		{
			name: "Coordinates equal.",
			coordinate1: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinate2: types.Coordinate{
				X: 1,
				Y: 0,
			},
			want: true,
		},
		{
			name: "Coordinates equal only in y direction.",
			coordinate1: types.Coordinate{
				X: 1,
				Y: 1,
			},
			coordinate2: types.Coordinate{
				X: 2,
				Y: 1,
			},
			want: false,
		},
		{
			name: "Coordinates don't equal.",
			coordinate1: types.Coordinate{
				X: 1,
				Y: 0,
			},
			coordinate2: types.Coordinate{
				X: -1,
				Y: 0,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.AreCoordinatesEqual(tt.coordinate1, tt.coordinate2)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestShiftCoordinates(t *testing.T) {
	tests := []struct {
		name        string
		x           int
		y           int
		coordinates []types.Coordinate
		want        []types.Coordinate
	}{
		{
			name: "Coordinates shift.",
			x:    1,
			y:    2,
			coordinates: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
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
					X: -21,
					Y: 1,
				},
			},
			want: []types.Coordinate{
				{
					X: 1,
					Y: 2,
				},
				{
					X: 2,
					Y: 2,
				},
				{
					X: 1,
					Y: 3,
				},
				{
					X: 2,
					Y: 3,
				},
				{
					X: -20,
					Y: 3,
				},
			},
		},
		{
			name: "Coordinates shift with negative shift.",
			x:    -1,
			y:    -3,
			coordinates: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
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
			},
			want: []types.Coordinate{
				{
					X: -1,
					Y: -3,
				},
				{
					X: 0,
					Y: -3,
				},
				{
					X: -1,
					Y: -2,
				},
				{
					X: 0,
					Y: -2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ShiftCoordinates(tt.x, tt.y, tt.coordinates)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}

func TestShiftCoordinate(t *testing.T) {
	tests := []struct {
		name       string
		x          int
		y          int
		coordinate types.Coordinate
		want       types.Coordinate
	}{
		{
			name: "Coordinate shift.",
			x:    1,
			y:    2,
			coordinate: types.Coordinate{
				X: 1,
				Y: 3,
			},
			want: types.Coordinate{
				X: 2,
				Y: 5,
			},
		},
		{
			name: "Coordinate shift in negative direction.",
			x:    -4,
			y:    -2,
			coordinate: types.Coordinate{
				X: 2,
				Y: 1,
			},
			want: types.Coordinate{
				X: -2,
				Y: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ShiftCoordinate(tt.x, tt.y, tt.coordinate)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("The test '%s' did not return the expected results.\nExpected:\n%v\nActual:\n%v", tt.name, tt.want, result)
			}
		})
	}
}
