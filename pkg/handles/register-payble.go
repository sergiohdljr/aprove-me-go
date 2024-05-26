package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handle_errors "github.com/sergiohdljr/aprove-me-go/pkg/handles/errors"
)

type PaybleRequest struct {
	ID              string          `json:"id"  binding:"required,uuid"`
	Value           float64         `json:"value" binding:"required" `
	EmissionDate    string          `json:"emissionDate" binding:"required" `
	Assignor        string          `json:"assignor" binding:"required" `
	AssignorDetails AssignorDetails `json:"assignorDetails" binding:"required" `
}

type AssignorDetails struct {
	ID       string `json:"id" binding:"required"`
	Document string `json:"document" binding:"required,min=11,max=30"`
	Email    string `json:"email" binding:"required,email,max=140"`
	Phone    string `json:"phone" binding:"required,max=20"`
	Name     string `json:"name" binding:"required,max=140"`
}

func RegisterPayble(ctx *gin.Context) {
	var body PaybleRequest

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, handle_errors.ValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, body)
}
