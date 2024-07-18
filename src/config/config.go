package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Cors     CorsConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	OTP      OTPConfig
}

type ServerConfig struct {
	Port    string
	runMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Type     string
}

type CorsConfig struct {
	AllowOrigins string
}

type PostgresConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DBName             string
	SSLMode            string
	MaxIdleConnections int
	MaxOpenConnections int
	MaxLifetime        time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type PasswordConfig struct {
	IncludeChars  bool
	IncludeDigits bool
	IncludeUppers bool
	IncludeLowers bool
	MinLength     int
	MaxLength     int
}

type OTPConfig struct {
	Digits      int
	ExpiresTime time.Duration
	Limiter     time.Duration
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(configPath, "yaml")
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("unable to parse config: %v", err)
	}
	return config
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to parse config: %v", err)
		return nil, err
	}
	return &config, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	configPaths := map[string]string{
		"development": "/config/config-development",
		"production":  "/config/config-production",
		"docker":      "/config/config-docker",
	}

	path, exists := configPaths[env]
	if !exists {
		return configPaths["development"]
	}
	return path
}
