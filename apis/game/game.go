package game

import (
	"net/http"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	basePath := "./images/pieces"
	var details []types.PieceDetails = utils.CreateAllPieceDetails(basePath)
	c.IndentedJSON(http.StatusOK, details)
}
