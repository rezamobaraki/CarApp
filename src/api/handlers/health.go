package handlers

import (
	"github.com/rezamobaraki/CarApp/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health godoc
// @Summary Health check
// @Description Check the health of the service
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed to validate the request"
// @Failure 500 {object} helper.BaseHttpResponse{Error=[]error} "internal server error"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(context *gin.Context) {
	context.JSON(200, helper.GenerateBaseResponse("Boom Boom 💥", true, 0))
	return
}
