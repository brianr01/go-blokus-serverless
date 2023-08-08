package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/brianr01/go-blockus-serverless/apis/game"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g.GET("/pong", func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})

	gameRoutes := g.Group("/game")
	gameRoutes.GET("/pieces", game.GetAllPieces)
	gameRoutes.GET("/moves", game.GetValidMoves)
	gameRoutes.POST("/transform-grid", game.TransformGrid)

	env := os.Getenv("GIN_MODE")

	if env == "release" {
		ginLambda = ginadapter.New(g)

		lambda.Start(Handler)
	} else {
		g.Run(":8080")
	}

	ginLambda = ginadapter.New(g)
	lambda.Start(Handler)
}
