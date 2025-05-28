package main

import (
	"file-manager/db"
	"file-manager/middleware"
	"file-manager/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	db.Init_Mis()

	e := echo.New()

	// configureMiddleware(e)
	middleware.ConfigureMiddleware(e)
	routes.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
