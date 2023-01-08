package main

import (
	"fmt"
	"log"
	"os"

	"absensi/configs"
	"absensi/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	//load database
	configs.InitDB()

	//load migration
	configs.InitMigrate()

	//init Router
	e := echo.New()
	routes.InitRouter(e)

	//start Echo
	e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
