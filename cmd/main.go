package main

import "github.com/Ahmad-mufied/go-digilib/config"

func main() {
	config.InitViper()
	//postgresDb := config.InitDB()

	//dbModel := data.New(postgresDb)
	//handler.InitHandler(dbModel)
	//
	//e := echo.New()
	//
	//server.Routes(e)
	//
	//// Memulai server pada port 8080
	//log.Println("Starting Server")
	//
	//env := config.Viper.GetString("APP_ENV")
	//port := "8080"
	//
	//if env == "production" {
	//	log.Println("Running in production mode")
	//	port = config.Viper.GetString("PORT")
	//} else {
	//	log.Println("Running in development mode")
	//}
	//
	//e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}
