package services

import (
	"fmt"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/rezamobaraki/CarApp/constants"
	"github.com/rezamobaraki/CarApp/data/cache"
	"github.com/rezamobaraki/CarApp/pkg/logging"
	"github.com/rezamobaraki/CarApp/pkg/service_errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type OTPService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OTPDto struct {
	Value string
	Used  bool
}

func NewOTPService(cfg *config.Config) *OTPService {
	logger := logging.NewLogger(cfg)
	redisClient := cache.GetRedis()
	return &OTPService{logger: logger, cfg: cfg, redisClient: redisClient}
}

func (s *OTPService) SetOTP(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OTPDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OTPDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPUsed}
	}
	fmt.Println(s.cfg.OTP.ExpireTime * time.Second)
	fmt.Println(s.cfg.OTP.ExpireTime)
	err = cache.Set(s.redisClient, key, val, s.cfg.OTP.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OTPService) ValidateOTP(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[OTPDto](s.redisClient, key)
	if err != nil {
		return err
	}
	if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTPUsed}
	}
	if res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OTONotValid}
	}
	res.Used = true
	return cache.Set(s.redisClient, key, res, s.cfg.OTP.ExpireTime*time.Second)
}
