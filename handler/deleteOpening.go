package handler

import (
	"fmt"
	"net/http"

	"github.com/AnielHill/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHangler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query").Error())
		return
	}
	opening := schemas.Opening{}
	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
	}
	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
