package handlers

import (
	"github.com/rezamobaraki/CarApp/api/dto"
	"github.com/rezamobaraki/CarApp/api/helper"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/rezamobaraki/CarApp/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{
		service: services.NewUserService(cfg),
	}
}

// SendOTP godoc
// @Summary Send OTP
// @Description Send OTP to user
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.GetUserRequest true "User details for OTP"
// @Success 201 {object} helper.BaseHttpResponse
// @Failure 400 {object} helper.BaseHttpResponse
// @Failure 409 {object} helper.BaseHttpResponse
// @Router /v1/users/send-otp/ [post]
func (h *UserHandler) SendOTP(context *gin.Context) {
	request := new(dto.GetUserRequest)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(
			nil, false, -1, err))
		return
	}

	err = h.service.SendOTP(request)
	if err != nil {
		context.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return
	}
	// Call internal SMS Service
	context.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
