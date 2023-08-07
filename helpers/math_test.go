package helpers_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/helpers"
)

func TestIsRotationValidFor90Degrees(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "Zero degrees is valid.",
			input: 0,
			want:  true,
		},
		{
			name:  "90 degrees is valid.",
			input: 90,
			want:  true,
		},
		{
			name:  "180 degrees is valid.",
			input: 270,
			want:  true,
		},
		{
			name:  "270 degrees is valid.",
			input: 270,
			want:  true,
		},
		{
			name:  "360 degrees is valid.",
			input: 360,
			want:  true,
		},
		{
			name:  "450 degrees is valid.",
			input: 450,
			want:  true,
		},
		{
			name:  "-90 degrees is valid.",
			input: -90,
			want:  true,
		},
		{
			name:  "1 degrees is invalid.",
			input: 1,
			want:  false,
		},
		{
			name:  "89 degrees is invalid.",
			input: 89,
			want:  false,
		},
		{
			name:  "91 degrees is invalid.",
			input: 91,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.IsRotationValidFor90Degrees(tt.input)

			if result != tt.want {
				t.Errorf(
					"Test '%s' expected result '%t' recieved '%t",
					tt.name,
					tt.want,
					result,
				)
			}
		})
	}
}

func TestTranspose2d(t *testing.T) {
	// Create a new directory for testing.
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "Transpose of identiy matrix is identity matrix.",
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			want: [][]int{
				{1, 0},
				{0, 1},
			},
		},
		{
			name: "Transpose of upper triangular matrix equals lower triangluar matrix.",
			input: [][]int{
				{1, 1},
				{0, 1},
			},
			want: [][]int{
				{1, 0},
				{1, 1},
			},
		},
		{
			name: "Transpose of row matrix is column matrix",
			input: [][]int{
				{1, 0},
			},
			want: [][]int{
				{1},
				{0},
			},
		},
		{
			name: "Transpose of column matrix is row matrix",
			input: [][]int{
				{1, 0, 1, 1, 0, 1},
			},
			want: [][]int{
				{1},
				{0},
				{1},
				{1},
				{0},
				{1},
			},
		},
		{
			name: "Transpose of rectangular matrix is rectangular",
			input: [][]int{
				{1, 1, 0, 0, 0, 1},
				{1, 0, 0, 1, 1, 1},
				{0, 1, 1, 0, 1, 1},
			},
			want: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 1},
				{0, 1, 0},
				{0, 1, 1},
				{1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.Transpose2d(tt.input)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Transpose operation '%s' did not retun the expected output.\n Want:\n%s\nResult:\n%s\n",
					tt.name, helpers.GetStringFrom2d(tt.want),
					helpers.GetStringFrom2d(result),
				)
			}
		})
	}
}

func TestReverseRows2d(t *testing.T) {
	// Create a new directory for testing.
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "Rows reverse on identity matrix.",
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "Rows reverse on upper triangular matrix.",
			input: [][]int{
				{1, 1, 1},
				{0, 1, 1},
				{0, 0, 1},
			},
			want: [][]int{
				{0, 0, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
		},
		{
			name: "Rows reverse on rectangular matrix.",
			input: [][]int{
				{1, 1, 1, 1, 0, 1},
				{0, 1, 1, 0, 1, 1},
			},
			want: [][]int{
				{0, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.ReverseRows2d(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Reverse rows 2d operation '%s' did not retun the expected output.\n Want:\n%s\nResult:\n%s\n",
					tt.name, helpers.GetStringFrom2d(tt.want),
					helpers.GetStringFrom2d(result),
				)
			}
		})
	}
}

func TestReverseColumns2d(t *testing.T) {
	// Create a new directory for testing.
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "Column reverse on identity matrix.",
			input: [][]int{
				{1, 0},
				{0, 1},
			},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "Column reverse on upper triangular matrix.",
			input: [][]int{
				{1, 1, 1},
				{0, 1, 1},
				{0, 0, 1},
			},
			want: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 0},
			},
		},
		{
			name: "Column reverse on rectangular matrix.",
			input: [][]int{
				{0, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 0, 1},
			},
			want: [][]int{
				{1, 1, 0, 1, 1, 0},
				{1, 0, 1, 1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.ReverseColumns2d(tt.input)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Reverse Column 2d operation '%s' did not retun the expected output.\n Want:\n%s\nResult:\n%s\n",
					tt.name, helpers.GetStringFrom2d(tt.want),
					helpers.GetStringFrom2d(result),
				)
			}
		})
	}
}
