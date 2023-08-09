package game

import (
	"log"
	"net/http"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

type GetBestMovesRequest struct {
	ColorNumber       types.ColorNumber
	Grid              types.Grid
	ProbabilityGrid   types.ProbabilityGrid
	AvailabilityGrid  types.AvailabilityGrid
	ProbabilityPieces []types.ProbabilityPiece
	PieceDetails      []types.PieceDetail
}

type BestMove struct {
	Move types.Move
	Grid types.Grid
}

func GetBestMoves(c *gin.Context) {
	var requestBody GetBestMovesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("An error has occured.")
	}

	g := requestBody.Grid
	pg := requestBody.ProbabilityGrid
	ag := requestBody.AvailabilityGrid
	pps := requestBody.ProbabilityPieces
	pds := requestBody.PieceDetails

	pg = utils.SetUnavailableCoordinatesToZero(ag, pg)
	clr := requestBody.ColorNumber

	csRanked := utils.GetHighestRankingCoordinates(pg)
	ppsIndexesRanked := utils.GetHighestRankingPieces(pps)

	pdsRanked := make([]types.PieceDetail, 0)
	for _, index := range ppsIndexesRanked {
		pdsRanked = append(pdsRanked, pds[index])
	}

	bestMoves := make([]BestMove, 0)
	for _, cRanked := range csRanked {
		for _, pdRanked := range pdsRanked {
			ms := utils.CreateMovesAtCoordinatesForPieceDetails([]types.PieceDetail{pdRanked}, []types.Coordinate{cRanked}, clr)
			msValid := utils.FilterMovesAllowedMoves(ms, ag, g)

			if len(msValid) > 0 {
				for _, mValid := range msValid {
					gridWithMValid := utils.GetGridWithValidMove(mValid, g)

					bestMoves = append(bestMoves, BestMove{
						Move: mValid,
						Grid: gridWithMValid,
					})
				}
			}
		}
	}

	c.IndentedJSON(http.StatusOK, bestMoves)
}
