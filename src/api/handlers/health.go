package handlers

import (
	"github.com/MrRezoo/CarApp/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(context *gin.Context) {
	context.JSON(200, helper.GenerateBaseResponse("Boom Boom 💥", true, 0))
	return
}
