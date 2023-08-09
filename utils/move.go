package utils

import (
	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
)

func CreateMovesAtCoordinatesForPieceDetails(pds []types.PieceDetail, cs []types.Coordinate, clr types.ColorNumber) []types.Move {
	ms := make([]types.Move, 0)

	for _, c := range cs {
		ms = append(ms, CreateMovesAtCoordinateForPieceDetails(pds, c, clr)...)
	}

	return ms
}

func CreateMovesAtCoordinateForPieceDetails(pds []types.PieceDetail, c types.Coordinate, clr types.ColorNumber) []types.Move {
	ms := make([]types.Move, 0)

	for _, pd := range pds {
		ms = append(ms, CreateMovesAtCoordinateForPieceDetail(pd, c, clr)...)
	}

	return ms
}

func CreateMovesAtCoordinateForPieceDetail(pd types.PieceDetail, c types.Coordinate, clr types.ColorNumber) []types.Move {
	ms := make([]types.Move, 0)

	for _, s := range pd.Symmetries {
		ms = append(ms, CreateMovesAtCoordinateForSymmetry(s, c, clr)...)
	}

	return ms
}

func CreateMovesAtCoordinateForSymmetry(s types.Symmetry, c types.Coordinate, clr types.ColorNumber) []types.Move {
	cs := ShiftCoordinatesByCoordinate(c, s.PlayableCoordinates)

	return CreateMovesAtCoordinates(s.RidgidPiece, clr, cs)
}

func CreateMovesAtCoordinates(rp types.RidgidPiece, clr types.ColorNumber, cs []types.Coordinate) []types.Move {
	ms := make([]types.Move, len(cs))

	for i, c := range cs {
		ms[i] = types.Move{
			Coordinate:  c,
			ColorNumber: clr,
			RidgidPiece: rp,
		}
	}

	return ms
}

func FilterMovesAllowedMoves(ms []types.Move, ag types.AvailabilityGrid, g types.Grid) []types.Move {
	msValid := make([]types.Move, 0)
	for _, m := range ms {
		if IsMoveAllowed(m, ag, g) {
			msValid = append(msValid, m)
		}
	}

	return msValid
}

func IsMoveAllowed(m types.Move, ag types.AvailabilityGrid, g types.Grid) bool {
	cs := GetColorCoordinatesForMove(m)

	if !AreCoordinatesOnGrid(cs, g) {
		return false
	}

	ans := GetAvailabilityNumbersFromCoordinates(cs, ag)

	// There is no playable location in the locations.
	if !AvailabilityNumberInAvailabilityNumbers(constants.AvailabilityNumberPlayable, ans) {
		return false
	}

	// There is a location that is not at least valid.
	if !AvailabilityNumbersWithMinium(constants.AvailabilityNumberValid, ans) {
		return false
	}

	return true
}

func GetColorCoordinatesForMove(m types.Move) []types.Coordinate {
	cs := make([]types.Coordinate, 0)
	rpCoords := GetColorCoordinatesForRidgidPiece(m.RidgidPiece)
	for _, rpCoord := range rpCoords {
		cs = append(cs, AddCoordinates(m.Coordinate, rpCoord))
	}

	return cs
}
