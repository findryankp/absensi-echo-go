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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	e.Start(fmt.Sprintf(":%v", port))
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
