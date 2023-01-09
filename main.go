package main

import (
	"absensi/configs"
	"absensi/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//load Configuration
	loadConfiguration()

	//init Router
	e := echo.New()
	routes.InitRouter(e)

	//start Echo
	e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

func loadConfiguration() {
	//load documentation
	configs.InitSwagger()

	//Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	//load database
	configs.InitDB()

	//load migration
	configs.InitMigrate()
}
