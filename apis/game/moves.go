package game

import (
	"log"
	"net/http"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

type GetValidMovesRequest struct {
	ColorNumber  types.ColorNumber
	Grid         types.Grid
	PieceDetails []types.PieceDetail
}

type ValidMovesResponse struct {
	Move types.Move
	Grid types.Grid
}

func GetValidMoves(c *gin.Context) {
	var requestBody GetBestMovesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("An error has occured.")
	}

	g := requestBody.Grid
	clr := requestBody.ColorNumber

	pds := requestBody.PieceDetails

	ag := utils.GetAvailabilityGridFromGrid(clr, g)
	cs := utils.GetCoordiantesWithMinAvailabilityNumber(constants.AvailabilityNumberPlayable, ag)

	ms := utils.CreateMovesAtCoordinatesForPieceDetails(pds, cs, clr)

	res := make([]ValidMovesResponse, 0)

	if len(ms) > 0 {
		for _, m := range ms {
			gWithM := utils.GetGridWithValidMove(m, g)

			res = append(res, ValidMovesResponse{
				Move: m,
				Grid: gWithM,
			})
		}
	}

	c.IndentedJSON(http.StatusOK, len(ms))
}
