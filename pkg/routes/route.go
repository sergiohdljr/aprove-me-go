package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiohdljr/aprove-me-go/pkg/handles"
)

func RouteInnit() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	V1 := r.Group("/api/v1")
	{
		V1.POST("/integrations/payable", handles.RegisterPayble)
	}

	return r
}
