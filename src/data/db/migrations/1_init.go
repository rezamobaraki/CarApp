package migrations

import (
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/db"
	"github.com/MrRezoo/CarApp/data/models"
	"github.com/MrRezoo/CarApp/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDB()

	var tables []interface{}

	country := models.Country{}
	city := models.City{}
	if !database.Migrator().HasTable(country) {
		tables = append(tables, country, city)
	}
	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables created", nil)

}

func Rollback() {

}
