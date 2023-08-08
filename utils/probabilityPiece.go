package utils

import "github.com/brianr01/go-blockus-serverless/types"

func GetHighestRankingPieces(pps []types.ProbabilityPiece) []int {
	var highestRank types.ProbabilityPiece
	var highestRankingPiece int
	for x := 0; x < len(pps); x++ {
		rank := pps[x]
		if highestRank < rank {
			highestRank = rank
			highestRankingPiece = x
		}
	}

	return []int{
		highestRankingPiece,
	}
}
