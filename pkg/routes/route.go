package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	ID              string          `json:"id"`
	Value           float64         `json:"value"`
	EmissionDate    string          `json:"emissionDate"`
	Assignor        string          `json:"assignor"`
	AssignorDetails AssignorDetails `json:"assignorDetails"`
}

type AssignorDetails struct {
	ID       string `json:"id"`
	Document string `json:"document"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

func RouteInnit() *gin.Engine {
	// var db = new(map[string]string)

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/integrations/payable", func(ctx *gin.Context) {
		var body PaymentRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"body": body,
		})
	})

	return r
}
