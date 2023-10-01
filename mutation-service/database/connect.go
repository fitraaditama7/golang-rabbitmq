package database

import (
	"fmt"
	"mutation-service/config"
	"mutation-service/utils/logger"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Config("DB_HOST"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		config.Config("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log := logger.Log()
		log.Error().Err(err).Msg(err.Error())
		return nil
	}
	return db
}
