package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiohdljr/aprove-me-go/pkg/database"
	"github.com/sergiohdljr/aprove-me-go/pkg/models"
	"net/http"
)

func GetPayble(ctx *gin.Context) {
	var payble models.Payment
	database.Db.First(&payble, "id = ?", ctx.Param("id"))
	ctx.JSON(http.StatusOK, payble)
}
