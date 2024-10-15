package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nostalgia-be/internal/app"
	"nostalgia-be/internal/database"
	"os"
)

func main() {
	e := echo.New()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		e.Logger.Fatal("Error loading .env file:", err)
	}

	// Connect to the database
	database.Connect()
	defer database.Close()

	// Initialize the Application struct
	application := app.New(e)

	// Register middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	application.RegisterRoutes(e)

	// Start the server
	port := os.Getenv("APP_PORT")
	appAddress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAddress))
}
