package game

import (
	"log"
	"net/http"

	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

type AvailabilityGridRequest struct {
	ColorNumber types.ColorNumber
	Grid        types.Grid
}

func GetAvailabilityGrid(c *gin.Context) {
	var requestBody AvailabilityGridRequest

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("An error has occured.")
	}

	g := requestBody.Grid
	clr := requestBody.ColorNumber

	ag := utils.GetAvailabilityGridFromGrid(clr, g)

	c.IndentedJSON(http.StatusOK, ag)
}
