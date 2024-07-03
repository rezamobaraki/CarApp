package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(context *gin.Context) {
	context.JSON(200, "Boom Boom 💥")
	return
}
