package main

import (
	"fmt"
	"github.com/Ahmad-mufied/go-digilib/config"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/server"
	"github.com/Ahmad-mufied/go-digilib/server/handler"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	config.InitViper()
	postgresDb := config.InitDB()

	dbModel := data.New(postgresDb)
	handler.InitHandler(dbModel)

	e := echo.New()

	server.Routes(e)

	// Memulai server pada port 8080
	log.Println("Starting Server")

	env := config.Viper.GetString("APP_ENV")
	port := "8080"

	if env == "production" {
		log.Println("Running in production mode")
		port = config.Viper.GetString("PORT")
	} else {
		log.Println("Running in development mode")
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}
