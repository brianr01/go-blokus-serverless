package utils

import (
	"reflect"

	"github.com/brianr01/go-blockus-serverless/types"
)

func AddCoordinates(cs ...types.Coordinate) types.Coordinate {
	x := 0
	y := 0

	for _, c := range cs {
		x += c.X
		y += c.Y
	}

	return types.Coordinate{
		X: x,
		Y: y,
	}
}

func FilterCoordinatesByMatchingCoordinates(csToFilter []types.Coordinate, csToMatch []types.Coordinate) []types.Coordinate {
	csFiltered := make([]types.Coordinate, 0)

	for _, c := range csToFilter {
		if IsCoordinateInCoordinates(c, csToMatch) {
			csFiltered = append(csFiltered, c)
		}
	}

	return csFiltered
}

func IsCoordinateInCoordinates(c types.Coordinate, cs []types.Coordinate) bool {
	for _, cToCompare := range cs {
		if AreCoordinatesEqual(c, cToCompare) {
			return true
		}
	}

	return false
}

func ShiftCoordinatesByCoordinate(c types.Coordinate, cs []types.Coordinate) []types.Coordinate {
	return ShiftCoordinates(c.X, c.Y, cs)
}

func AreCoordinatesEqual(c1 types.Coordinate, c2 types.Coordinate) bool {
	return reflect.DeepEqual(c1, c2)
}

func ShiftCoordinates(x int, y int, cs []types.Coordinate) []types.Coordinate {
	csResult := make([]types.Coordinate, len(cs))

	for i, c := range cs {
		csResult[i] = ShiftCoordinate(x, y, c)
	}

	return csResult
}

func ShiftCoordinate(x int, y int, c types.Coordinate) types.Coordinate {
	return types.Coordinate{
		X: c.X + x,
		Y: c.Y + y,
	}
}
