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
	var highestRank float64
	var highestRankingCoord types.Coordinate

	for x := 0; x < len(pg); x++ {
		for y := 0; y < len(pg[0]); y++ {
			rank := pg[x][y]
			if highestRank < rank {
				highestRank = rank
				highestRankingCoord = types.Coordinate{
					X: x,
					Y: y,
				}
			}
		}
	}

	return []types.Coordinate{
		highestRankingCoord,
	}
}
