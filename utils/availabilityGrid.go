package utils

import (
	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
)

// // constants.AvailabilityNumberPlayable, clr, g
// func GetCoordinatesWithMinAvailabilityNumber(an types.AvailabilityNumber, clr types.ColorNumber, g types.Grid ) {

// }

func GetAvailabilityNumbersFromCoordinates(cs []types.Coordinate, ag types.AvailabilityGrid) []types.AvailabilityNumber {
	as := make([]types.AvailabilityNumber, 0)
	for _, c := range cs {
		as = append(as, getAvailabilityNumberForCoordinate(c, ag))
	}

	return as
}

func getAvailabilityNumberForCoordinate(c types.Coordinate, ag types.AvailabilityGrid) types.AvailabilityNumber {
	return ag[c.X][c.Y]
}

func GetCoordiantesWithMinAvailabilityNumber(anMin types.AvailabilityNumber, ag types.AvailabilityGrid) []types.Coordinate {
	cs := make([]types.Coordinate, 0)
	agSize := GetAvailabilityGridSize(ag)
	w, h := agSize[0], agSize[1]

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if ag[x][y] >= anMin {
				cs = append(cs, types.Coordinate{
					X: x,
					Y: y,
				})
			}
		}
	}

	return cs
}

func GetAvailabilityGridFromGrid(clr types.ColorNumber, g types.Grid) types.AvailabilityGrid {
	gSize := GetGridSize(g)
	w, h := gSize[0], gSize[1]
	a := CreateEmptyAvailabilityGrid(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := types.Coordinate{
				X: x,
				Y: y,
			}

			a[x][y] = CreateAvailabilityForCoordinate(c, clr, g)
		}
	}

	return a
}

func CreateAvailabilityForCoordinate(c types.Coordinate, clr types.ColorNumber, g types.Grid) types.AvailabilityNumber {
	// If the coord is occupied.
	if !IsCoordinateOnGridEmpty(c, g) {
		return constants.AvailabilityNumberOccupied
	}

	// If a edge color matches.
	if DoesEdgesToCoordnateHaveColor(c, clr, g) {
		return constants.AvailabilityNumberEmpty
	}

	// If the piece is not playable.
	if !IsPlayableCoordinate(c, clr, g) {
		return constants.AvailabilityNumberValid
	}

	return constants.AvailabilityNumberPlayable
}

func CreateEmptyAvailabilityGrid(w int, h int) types.AvailabilityGrid {
	var a types.AvailabilityGrid

	for i := 0; i < w; i++ {
		var availabilityRow []types.AvailabilityNumber
		for j := 0; j < h; j++ {
			availabilityRow = append(availabilityRow, constants.AvailabilityNumberEmpty)
		}
		a = append(a, availabilityRow)
	}

	return a
}

func GetAvailabilityGridSize(g types.AvailabilityGrid) [2]int {
	return [2]int{len(g), len(g[0])}
}
