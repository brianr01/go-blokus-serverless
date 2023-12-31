package utils_test

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestGetColorCoordinatesForRidgidPiece(t *testing.T) {
	tests := []struct {
		name        string
		ridgidPiece types.RidgidPiece
		want        []types.Coordinate
	}{
		{
			name: "one piece",
			ridgidPiece: types.RidgidPiece{
				{
					1,
				},
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
			},
		},
		{
			name: "triomino L",
			ridgidPiece: types.RidgidPiece{
				{1, 0},
				{1, 0},
				{1, 1},
			},
			want: []types.Coordinate{
				{
					X: 0,
					Y: 0,
				},
				{
					X: 1,
					Y: 0,
				},
				{
					X: 2,
					Y: 0,
				},
				{
					X: 2,
					Y: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetColorCoordinatesForRidgidPiece(tt.ridgidPiece)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("The result %v does not match %v", result, tt.want)
			}
		})
	}
}

func TestRotateRidgidPiece(t *testing.T) {
	tests := []struct {
		name        string
		ridgidPiece types.RidgidPiece
		rotation    int
		want        types.RidgidPiece
	}{
		{
			name: "Rotate 90 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: 90,
			want: types.RidgidPiece{
				{0, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "Rotate 180 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: 180,
			want: types.RidgidPiece{
				{0, 1, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
		{
			name: "Rotate 270 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: 270,
			want: types.RidgidPiece{
				{0, 1, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		{
			name: "Rotate 360 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: 360,
			want: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "Rotate 450 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: 450,
			want: types.RidgidPiece{
				{0, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "Rotate -90 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: -90,
			want: types.RidgidPiece{
				{0, 1, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		{
			name: "Rotate -180 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: -180,
			want: types.RidgidPiece{
				{0, 1, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
		{
			name: "Rotate -270 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			rotation: -270,
			want: types.RidgidPiece{
				{0, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "Rotate rectangle matrix 90 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 0, 0},
				{1, 1, 1},
			},
			rotation: 90,
			want: types.RidgidPiece{
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		{
			name: "Rotate other rectangle matrix 90 degrees.",
			ridgidPiece: types.RidgidPiece{
				{0, 1},
				{0, 1},
				{1, 1},
			},
			rotation: 90,
			want: types.RidgidPiece{
				{1, 0, 0},
				{1, 1, 1},
			},
		},
		{
			name: "Rotate other rectangle matrix -90 degrees.",
			ridgidPiece: types.RidgidPiece{
				{1, 1},
				{1, 0},
				{1, 0},
			},
			rotation: -90,
			want: types.RidgidPiece{
				{1, 0, 0},
				{1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.RotateRidgidPiece(tt.ridgidPiece, tt.rotation)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Test '%s' did not recieve the expectecd value.\nExpected:\n%s\nActual:\n%s\n",
					tt.name,
					utils.GetStringFrom2d(tt.want),
					utils.GetStringFrom2d(result),
				)
			}
		})
	}
}

func TestCreateRidigPieceFromImage(t *testing.T) {
	w := color.RGBA{255, 255, 255, 0xff}
	b := color.RGBA{0, 0, 0, 0xff}
	tests := []struct {
		name  string
		input [][]color.RGBA
		want  types.RidgidPiece
	}{
		{
			name: "OnePiece",
			input: [][]color.RGBA{
				{b},
			},
			want: types.RidgidPiece{
				{1},
			},
		},
		{
			name: "TrominoLPiece",
			input: [][]color.RGBA{
				{b, w},
				{b, b},
			},
			want: types.RidgidPiece{
				{1, 0},
				{1, 1},
			},
		},
		{
			name: "PentominoF",
			input: [][]color.RGBA{
				{b, b, w},
				{w, b, b},
				{w, b, w},
			},
			want: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "One Row",
			input: [][]color.RGBA{
				{b, b, w, b, w, w, w, b},
			},
			want: types.RidgidPiece{
				{1, 1, 0, 1, 0, 0, 0, 1},
			},
		},
		{
			name: "One Column",
			input: [][]color.RGBA{
				{b},
				{b},
				{w},
				{b},
				{w},
				{w},
				{w},
				{b},
			},
			want: types.RidgidPiece{
				{1},
				{1},
				{0},
				{1},
				{0},
				{0},
				{0},
				{1},
			},
		},
		{
			name: "many colors",
			input: [][]color.RGBA{
				{b, b, b, w, w, b, b, w},
				{w, b, b, w, w, b, b, w},
				{w, w, w, b, w, b, b, b},
				{b, w, b, b, w, b, w, w},
				{b, w, b, b, w, b, b, b},
				{w, w, w, w, w, b, b, w},
				{b, w, b, w, w, b, w, b},
				{w, b, w, w, w, b, b, b},
			},
			want: types.RidgidPiece{
				{1, 1, 1, 0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0, 1, 1, 1},
				{1, 0, 1, 1, 0, 1, 0, 0},
				{1, 0, 1, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 0, 1, 0, 1},
				{0, 1, 0, 0, 0, 1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := utils.CreateImageFrom2dColor(tt.input)
			result := utils.CreateRidigPieceFromImage(img)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Test '%s' did not recieve the expectecd value.\nExpected:\n%s\nActual:\n%s\n",
					tt.name,
					utils.GetStringFrom2d(tt.want),
					utils.GetStringFrom2d(result),
				)
			}
		})
	}
}

func TestGetDimensionsFromRidigPiece(t *testing.T) {
	tests := []struct {
		name  string
		input types.RidgidPiece
		want  [2]int
	}{
		{
			name: "OnePiece",
			input: types.RidgidPiece{
				{1},
			},
			want: [2]int{1, 1},
		},
		{
			name: "Many Columns",
			input: types.RidgidPiece{
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1},
			},
			want: [2]int{1, 16},
		},
		{
			name: "Many Rows",
			input: types.RidgidPiece{
				{1},
				{0},
				{1},
				{0},
				{1},
				{1},
				{1},
				{0},
				{1},
				{1},
				{0},
				{1},
				{1},
				{0},
				{1},
				{1},
			},
			want: [2]int{16, 1},
		},
		{
			name: "Many Columns And Rows",
			input: types.RidgidPiece{
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
			},
			want: [2]int{14, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetDimensionsFromRidigPiece(tt.input)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Test '%s' did not recieve expected value.\nExpected: %v\nActual:   %v\n",
					tt.name,
					tt.want,
					result,
				)
			}
		})
	}
}

func TestRidgidPieceInRidgidPieces(t *testing.T) {
	tests := []struct {
		name         string
		ridgidPiece  types.RidgidPiece
		ridgidPieces []types.RidgidPiece
		want         bool
	}{
		{
			name: "Matches with one item.",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{1},
				},
			},
			want: true,
		},
		{
			name: "Matches First Item",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{1},
				},
				{
					{0},
				},
				{
					{0},
				},
			},
			want: true,
		},
		{
			name: "Matches Middle Item",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{0},
				},
				{
					{1},
				},
				{
					{0},
				},
			},
			want: true,
		},
		{
			name: "Matches Last Item",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{0},
				},
				{
					{0},
				},
				{
					{1},
				},
			},
			want: true,
		},
		{
			name: "Doesn't Match",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{0},
				},
				{
					{0},
				},
				{
					{0},
				},
			},
			want: false,
		},
		{
			name: "Doesn't Match with similar",
			ridgidPiece: types.RidgidPiece{
				{1},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{1, 1},
				},
				{
					{1},
					{1},
				},
				{
					{1, 1},
					{1, 1},
				},
			},
			want: false,
		},
		{
			name: "Matches with complex",
			ridgidPiece: types.RidgidPiece{
				{1, 5, 1, 2, 3, 4},
				{5, 1, 2, 3, 3, 5},
				{3, 2, 1, 5, 3, 3},
				{2, 5, 4, 1, 1, 2},
			},
			ridgidPieces: []types.RidgidPiece{
				{
					{5, 5, 5, 2, 5, 5},
					{5, 5, 5, 5, 5, 5},
					{3, 5, 5, 5, 5, 5},
					{2, 5, 5, 5, 5, 5},
				},
				{
					{1},
					{1},
				},
				{
					{1, 5, 1, 2, 3, 4},
					{5, 1, 2, 3, 3, 5},
					{3, 2, 1, 5, 3, 3},
					{2, 5, 4, 1, 1, 2},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.RidgidPieceInRidgidPieces(tt.ridgidPiece, tt.ridgidPieces)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Test '%s' did not recieve expected value.\nExpected: %v\nActual:   %v\n",
					tt.name,
					tt.want,
					result,
				)
			}
		})
	}
}

func TestGetValidSymmetriesFromRidgidPiece(t *testing.T) {
	tests := []struct {
		name                      string
		input                     types.RidgidPiece
		ignorePlayableCoordinates bool
		want                      []types.Symmetry
	}{
		{
			name: "One Piece Symmetries",
			input: types.RidgidPiece{
				{1},
			},
			ignorePlayableCoordinates: false,
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1},
					},
					PlayableCoordinates: []types.Coordinate{
						{
							X: 0,
							Y: 0,
						},
					},
				},
			},
		},
		{
			name: "Tromino L Piece Symmetries",
			input: types.RidgidPiece{
				{1, 0},
				{1, 1},
			},
			ignorePlayableCoordinates: false,
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1, 0},
						{1, 1},
					},
					PlayableCoordinates: []types.Coordinate{
						{
							X: 0,
							Y: 0,
						},
						{
							X: 1,
							Y: 0,
						},
						{
							X: 1,
							Y: 1,
						},
					},
				},
				{
					Mirror:   true,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1, 1},
						{1, 0},
					},
					PlayableCoordinates: []types.Coordinate{
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
							Y: 0,
						},
					},
				},
				{
					Mirror:   false,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{1, 1},
						{0, 1},
					},
					PlayableCoordinates: []types.Coordinate{
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
					},
				},
				{
					Mirror:   true,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{0, 1},
						{1, 1},
					},
					PlayableCoordinates: []types.Coordinate{
						{
							X: 0,
							Y: 1,
						},
						{
							X: 1,
							Y: 0,
						},
						{
							X: 1,
							Y: 1,
						},
					},
				},
			},
		},
		{
			name: "Tromino I Piece Symmetries",
			input: types.RidgidPiece{
				{1},
				{1},
				{1},
			},
			ignorePlayableCoordinates: false,
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1},
						{1},
						{1},
					},
					PlayableCoordinates: []types.Coordinate{
						{
							X: 0,
							Y: 0,
						},
						{
							X: 2,
							Y: 0,
						},
					},
				},
				{
					Mirror:   false,
					Rotation: 1,
					RidgidPiece: types.RidgidPiece{
						{1, 1, 1},
					},
					PlayableCoordinates: []types.Coordinate{
						{
							X: 0,
							Y: 0,
						},
						{
							X: 0,
							Y: 2,
						},
					},
				},
			},
		},
		{
			name: "Penomino F Piece Symmetries",
			input: types.RidgidPiece{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			ignorePlayableCoordinates: true,
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1, 1, 0},
						{0, 1, 1},
						{0, 1, 0},
					},
				},
				{
					Mirror:   true,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{0, 1, 0},
						{0, 1, 1},
						{1, 1, 0},
					},
				},
				{
					Mirror:   false,
					Rotation: 1,
					RidgidPiece: types.RidgidPiece{
						{0, 0, 1},
						{1, 1, 1},
						{0, 1, 0},
					},
				},
				{
					Mirror:   true,
					Rotation: 1,
					RidgidPiece: types.RidgidPiece{
						{0, 1, 0},
						{1, 1, 1},
						{0, 0, 1},
					},
				},
				{
					Mirror:   false,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{0, 1, 0},
						{1, 1, 0},
						{0, 1, 1},
					},
				},
				{
					Mirror:   true,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{0, 1, 1},
						{1, 1, 0},
						{0, 1, 0},
					},
				},
				{
					Mirror:   false,
					Rotation: 3,
					RidgidPiece: types.RidgidPiece{
						{0, 1, 0},
						{1, 1, 1},
						{1, 0, 0},
					},
				},
				{
					Mirror:   true,
					Rotation: 3,
					RidgidPiece: types.RidgidPiece{
						{1, 0, 0},
						{1, 1, 1},
						{0, 1, 0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.GetValidSymmetriesFromRidgidPiece(tt.input)
			if tt.ignorePlayableCoordinates {
				for i := range result {
					symmetryResult := result[i]
					symmetry := tt.want[i]
					if !reflect.DeepEqual(symmetryResult.Mirror, symmetry.Mirror) {
						t.Errorf("The mirror value did mot match for symmetry with index %v", i)
					}

					if !reflect.DeepEqual(symmetryResult.RidgidPiece, symmetry.RidgidPiece) {
						t.Errorf(
							"The ridgnid peace symmetry with index %v.\nExpected: \n%v\nActual:   \n%v\n",
							i,
							symmetry.RidgidPiece,
							symmetryResult.RidgidPiece,
						)
					}

					if !reflect.DeepEqual(symmetryResult.Rotation, symmetry.Rotation) {
						t.Errorf(
							"The rotation value did not match for symmetry with index %v.\nExpected: \n%v\nActual:   \n%v\n",
							i,
							symmetry.Rotation,
							symmetryResult.Rotation,
						)
					}
				}

				return
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf(
					"Test '%s' did not recieve expected value.\nExpected: %v\nActual:   %v\n",
					tt.name,
					tt.want,
					result,
				)
			}
		})
	}
}
