package game

import (
	"log"
	"net/http"

	"github.com/brianr01/go-blockus-serverless/utils"
	"github.com/gin-gonic/gin"
)

type GetEmptyGridRequest struct {
	Width  int
	Height int
}

func GetEmptyGrid(c *gin.Context) {
	var requestBody GetEmptyGridRequest

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("An error has occured.")
	}

	w, h := requestBody.Width, requestBody.Height

	g := utils.CreateEmptyGrid(w, h)

	c.IndentedJSON(http.StatusOK, g)
}
