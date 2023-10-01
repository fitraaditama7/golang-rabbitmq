package config

import (
	"account-service/utils/logger"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log := logger.Log()
		log.Error().Err(err).Msg(err.Error())
	}
}

func Config(key string) string {
	return os.Getenv(key)
}

func GetEnvToDuration(e string) (d time.Duration, err error) {
	var envValue int
	envValue, err = strconv.Atoi(os.Getenv(e))
	d = time.Duration(envValue) * time.Second
	return
}
