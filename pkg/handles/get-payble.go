package handles

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sergiohdljr/aprove-me-go/pkg/database"
	"github.com/sergiohdljr/aprove-me-go/pkg/models"
	"net/http"
)

func GetPayble(ctx *gin.Context) {
	var payble models.Payment
	payble_id := ctx.Param("id")

	result := database.Db.First(&payble, "id = ?", payble_id)

	if result.Error != nil {
		ctx.Error(errors.New("Payble not found")).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, payble)
}
