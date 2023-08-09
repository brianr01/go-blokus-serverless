package utils

import (
	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
)

func SetUnavailableCoordinatesToZero(ag types.AvailabilityGrid, pg types.ProbabilityGrid) types.ProbabilityGrid {
	for x := 0; x < len(pg); x++ {
		for y := 0; y < len(pg[0]); y++ {
			if ag[x][y] < constants.AvailabilityNumberPlayable {
				pg[x][y] = 0
			}
		}
	}

	return pg
}

func GetHighestRankingCoordinates(pg types.ProbabilityGrid) []types.Coordinate {
	// TODO update to return best an not just non zero.
	cs := make([]types.Coordinate, 0)

	for x := 0; x < len(pg); x++ {
		for y := 0; y < len(pg[0]); y++ {
			rank := pg[x][y]
			if 0 < rank {
				cs = append(cs, types.Coordinate{
					X: x,
					Y: y,
				})
			}
		}
	}

	return cs
}
