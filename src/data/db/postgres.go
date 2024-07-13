package db

import (
	"fmt"
	"github.com/MrRezoo/CarApp/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbClient *gorm.DB

func InitDB(config *config.Config) error {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.DBName, config.Postgres.Password,
	)
	dbClient, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
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

	log.Printf("Postgres connected on %s:%s", config.Postgres.Host, config.Postgres.Port)
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	sqlDB, _ := dbClient.DB()
	_ = sqlDB.Close()
}
