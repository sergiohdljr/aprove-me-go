package handles

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sergiohdljr/aprove-me-go/pkg/database"
	handle_errors "github.com/sergiohdljr/aprove-me-go/pkg/handles/errors"
	"github.com/sergiohdljr/aprove-me-go/pkg/models"
)

type PaybleRequest struct {
	Value           float64         `json:"value" binding:"required" `
	EmissionDate    string          `json:"emissionDate" binding:"required" `
	Assignor        string          `json:"assignor" binding:"required" `
	AssignorDetails AssignorDetails `json:"assignorDetails" binding:"required" `
}

type AssignorDetails struct {
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

	assignor := models.Assignor{
		Document: body.AssignorDetails.Document,
		Email:    body.AssignorDetails.Email,
		Phone:    body.AssignorDetails.Phone,
		Name:     body.AssignorDetails.Name,
	}

	if err := database.Db.Create(&assignor).Error; err != nil {
		log.Fatalf("failed to create assignor: %v", err)
	}

	payment := models.Payment{
		Value:        body.Value,
		EmissionDate: time.Now(),
		AssignorID:   assignor.ID,
	}

	if err := database.Db.Create(&payment).Error; err != nil {
		log.Fatalf("failed to create payment: %v", err)
	}

	log.Printf("Payment created successfully with ID: %s", payment.ID)

	ctx.JSON(http.StatusOK, body)
}
