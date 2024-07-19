package services

import (
	"github.com/MrRezoo/CarApp/api/dto"
	"github.com/MrRezoo/CarApp/common"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/db"
	"github.com/MrRezoo/CarApp/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OTPService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		logger:     logging.NewLogger(cfg),
		cfg:        cfg,
		otpService: NewOTPService(cfg),
		database:   db.GetDB(),
	}
}

func (userService UserService) SendOTP(request *dto.GetUserRequest) error {
	otp := common.GenerateOtp()
	err := userService.otpService.SetOTP(request.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
