package game

import (
	"log"
	"net/http"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

type GridTransform struct {
	Grid      types.Grid
	ColorMap  map[types.ColorNumber]types.ColorNumber
	Rotations int
}

func TransformGrid(c *gin.Context) {
	var requestBody GridTransform

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("An error has occured.")
	}

	g := requestBody.Grid
	clrMap := requestBody.ColorMap
	rotations := requestBody.Rotations

	g = utils.RotateGrid(rotations*90, g)
	g = utils.SwitchColorNumbers(clrMap, g)

	c.IndentedJSON(http.StatusOK, g)
}
