package game

import (
	"net/http"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

func GetAllPieces(c *gin.Context) {
	var details []types.PieceDetails = utils.CreateAllPieceDetails(constants.PiecesImagePath)
	c.IndentedJSON(http.StatusOK, details)
}
