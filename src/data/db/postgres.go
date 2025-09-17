package db

import (
	"fmt"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/rezamobaraki/CarApp/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDB(config *config.Config) error {
	var err error
	err = createDatabase(&config.Postgres)
	if err != nil {
		return err
	}
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.DBName, config.Postgres.Password,
	)
	dbClient, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(config.Postgres.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.Postgres.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(config.Postgres.MaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.Startup, "Postgres connected", nil)
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	sqlDB, _ := dbClient.DB()
	_ = sqlDB.Close()
}

func createDatabase(cfg *config.PostgresConfig) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=%s dbname=postgres",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	var dbName string
	err = db.Raw("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower(?)", cfg.DBName).Scan(&dbName).Error
	if err != nil {
		return err
	}

	if dbName == "" {
		err = db.Exec("CREATE DATABASE " + cfg.DBName).Error
		if err != nil {
			return err
		}
		logger.Info(logging.Postgres, logging.Startup, "Database created", nil)
	} else {
		logger.Info(logging.Postgres, logging.Startup, "Database already exists", nil)
	}

	return nil
}
