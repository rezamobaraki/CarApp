package services

import (
	"github.com/rezamobaraki/CarApp/api/dto"
	"github.com/rezamobaraki/CarApp/common"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/rezamobaraki/CarApp/constants"
	"github.com/rezamobaraki/CarApp/data/db"
	"github.com/rezamobaraki/CarApp/data/models"
	"github.com/rezamobaraki/CarApp/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	cfg        *config.Config
	database   *gorm.DB
	logger     logging.Logger
	otpService *OTPService
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		cfg:        cfg,
		database:   db.GetDB(),
		logger:     logging.NewLogger(cfg),
		otpService: NewOTPService(cfg),
	}
}

func (userService *UserService) SendOTP(request *dto.GetUserRequest) error {
	otp := common.GenerateOtp()
	err := userService.otpService.SetOTP(request.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) getDefaultRole() (roleId int, err error) {

	if err = userService.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
