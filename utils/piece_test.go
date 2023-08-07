package utils_test

import (
	"reflect"
	"testing"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestGetValidSymmetriesFromRidgidPiece(t *testing.T) {
	tests := []struct {
		name  string
		input types.RidgidPiece
		want  []types.Symmetry
	}{
		{
			name: "One Piece Symmetries",
			input: types.RidgidPiece{
				{1},
			},
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1},
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
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1, 0},
						{1, 1},
					},
				},
				{
					Mirror:   true,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1, 1},
						{1, 0},
					},
				},
				{
					Mirror:   false,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{1, 1},
						{0, 1},
					},
				},
				{
					Mirror:   true,
					Rotation: 2,
					RidgidPiece: types.RidgidPiece{
						{0, 1},
						{1, 1},
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
			want: []types.Symmetry{
				{
					Mirror:   false,
					Rotation: 0,
					RidgidPiece: types.RidgidPiece{
						{1},
						{1},
						{1},
					},
				},
				{
					Mirror:   false,
					Rotation: 1,
					RidgidPiece: types.RidgidPiece{
						{1, 1, 1},
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
