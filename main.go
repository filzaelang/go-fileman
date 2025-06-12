package main

import (
	"file-manager/db"
	"file-manager/middleware"
	"file-manager/routes"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Logging to file
	logFile, err := os.OpenFile("D:/filemanager.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Tidak bisa membuka log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Load .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init_DB
	db.Init()
	db.Init_Mis()
	db.Init_Dev()

	e := echo.New()

	// Serve built assets
	e.Static("/assets", "dist/assets")
	e.Static("/js", "dist/js")
	e.Static("/", "dist")

	// configureMiddleware(e)
	middleware.ConfigureMiddleware(e)
	routes.ConfigureRoutes(e)

	// Delay when booting
	time.Sleep(5 * time.Second)

	// Run server
	if err := e.Start(":3000"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
