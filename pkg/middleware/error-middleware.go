package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		errors := ctx.Errors
		if len(errors) > 0 {
			for _, err := range errors {
				if err.Type == gin.ErrorTypePublic {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"status": "error",
						"error":  err.Error(),
					})
					return
				}
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "internal server error",
			})
		}

	}
}
