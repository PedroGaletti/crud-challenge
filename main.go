package main

import (
	"fmt"
	"os"

	"pismo/cmd/accounts"
	"pismo/cmd/transactions"
	"pismo/configs"
	"pismo/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	env = configs.Env
)

func main() {
	// * Configure *
	log.Logger = log.Output(zerolog.ConsoleWriter{TimeFormat: "02/01/2006 15:04:05", Out: os.Stderr})
	logLevel, _ := zerolog.ParseLevel(env.LogLevel)
	zerolog.SetGlobalLevel(logLevel)
	log.Info().Msg("Logs configured...")

	// * Database *
	log.Info().Msg("Configuring database...")
	database, err := db.InitMyqlDb(env.SqlUser, env.SqlPassword, env.SqlHost, env.SqlPort, env.SqlDb)
	if err != nil {
		log.Panic().Msg(fmt.Sprintf("Database connection error: %s", err))
	}

	log.Info().Msg("Running migrations...")
	db.Migrate(database)

	log.Info().Msg("Running Seeders...")
	db.Seed(database)

	// * Routes and Dependencies *
	log.Info().Msg("Creating routes and dependecies...")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	accounts.InjectDependency(router.Group("/accounts"), database)
	transactions.InjectDependency(router.Group("/transactions"), database)

	router.Run()
}
