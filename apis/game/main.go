package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func piece(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}
