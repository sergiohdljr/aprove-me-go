package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiohdljr/aprove-me-go/pkg/handles"
	"github.com/sergiohdljr/aprove-me-go/pkg/middleware"
)

func RouteInnit() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1/integrations")
	{
		api.Use(middleware.ErrorMiddleware())
		api.POST("/payble", handles.RegisterPayble)
		api.GET("/payble/:id", handles.GetPayble)

	}

	return r
}
