package services

import (
	"fmt"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/constants"
	"github.com/MrRezoo/CarApp/data/cache"
	"github.com/MrRezoo/CarApp/pkg/logging"
	"github.com/MrRezoo/CarApp/pkg/service_errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type OTPService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OTPDto struct {
	Value string `json:"value"`
	Used  bool   `json:"used"`
}

func NewOTPService(cfg *config.Config) *OTPService {
	logger := logging.NewLogger(cfg)
	redisClient := cache.GetRedis()
	return &OTPService{
		logger:      logger,
		cfg:         cfg,
		redisClient: redisClient,
	}
}

func (s *OTPService) SetOTP(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OTPDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get(s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OptExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}
	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OTPService) ValidateOTP(mobileNumber, otp string) error {
	return
}
