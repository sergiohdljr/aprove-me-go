package handles

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sergiohdljr/aprove-me-go/pkg/database"
	handle_errors "github.com/sergiohdljr/aprove-me-go/pkg/handles/errors"
	"github.com/sergiohdljr/aprove-me-go/pkg/models"
)

type PaybleRequest struct {
	Value           float64         `json:"value" binding:"required" `
	EmissionDate    string          `json:"emissionDate" binding:"required" `
	AssignorDetails AssignorDetails `json:"assignorDetails" binding:"required" `
}

type AssignorDetails struct {
	Document string `json:"document" binding:"required,min=11,max=30"`
	Email    string `json:"email" binding:"required,email,max=140"`
	Phone    string `json:"phone" binding:"required,max=20"`
	Name     string `json:"name" binding:"required,max=140"`
}

type PaybleResponse struct {
	ID           uuid.UUID        `json:"id"`
	Value        float64          `json:"value"`
	EmissionDate string           `json:"emissionDate"`
	AssinorID    uuid.UUID        `json:"assignor_id"`
	Assignor     AssignorResponse `json:"assignor"`
}

type AssignorResponse struct {
	ID       uuid.UUID `json:"id"`
	Document string    `json:"document"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Name     string    `json:"name"`
}

func RegisterPayble(ctx *gin.Context) {
	var body PaybleRequest

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, handle_errors.ValidationError(err))
		return
	}

	var assignor models.Assignor
	var payble models.Payment

	assignorExists := database.Db.First(&assignor, "document = ?", body.AssignorDetails.Document)

	if assignorExists.Error != nil {
		assignor = models.Assignor{
			Document: body.AssignorDetails.Document,
			Email:    body.AssignorDetails.Email,
			Phone:    body.AssignorDetails.Phone,
			Name:     body.AssignorDetails.Name,
		}

		if err := database.Db.Create(&assignor).Error; err != nil {
			log.Fatalf("failed to create assignor: %v", err)
		}
	}

	payble = models.Payment{
		Value:        body.Value,
		EmissionDate: time.Now(),
		AssignorID:   assignor.ID,
	}

	newPayment := database.Db.Create(&payble)

	if newPayment.Error != nil {
		ctx.Error(errors.New("failed to create payble"))
		return
	}

	response := PaybleResponse{
		ID:           payble.ID,
		Value:        payble.Value,
		EmissionDate: payble.EmissionDate.Format(time.RFC3339),
		AssinorID:    assignor.ID,
		Assignor: AssignorResponse{
			ID:       assignor.ID,
			Document: assignor.Document,
			Email:    assignor.Email,
			Phone:    assignor.Phone,
			Name:     assignor.Name,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
