package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	// Env export environment variables
	Env = GetConfig()
)

// Config : config data struct
type Config struct {
	GinMode     string
	LogLevel    string
	SqlDb       string
	SqlHost     string
	SqlPassword string
	SqlPort     string
	SqlUser     string
}

// GetConfig : get values of environment variables
func GetConfig() Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		log.Info().Msg("Runnning the application without a .env file.")
	}

	config.GinMode = os.Getenv("GIN_MODE")

	config.LogLevel = os.Getenv("LOG_LEVEL")

	config.SqlDb = os.Getenv("SQL_DB")
	config.SqlHost = os.Getenv("SQL_HOST")
	config.SqlPassword = os.Getenv("SQL_PASSWORD")
	config.SqlPort = os.Getenv("SQL_PORT")
	config.SqlUser = os.Getenv("SQL_USER")

	return config
}
