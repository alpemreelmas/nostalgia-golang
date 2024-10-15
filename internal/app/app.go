package app

import (
	"github.com/labstack/echo/v4"
	"nostalgia-be/internal/database"
	"nostalgia-be/internal/handlers"
)

// Application holds the dependencies for the application
type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler *handlers.Handler
}

// New creates a new Application instance
func New(echo *echo.Echo) *Application {

	// Initialize the database connection
	database.Connect()

	return &Application{
		logger:  echo.Logger,
		server:  echo,
		handler: &handlers.Handler{}, // Initialize your handler here
	}
}

// RegisterRoutes registers all application routes
func (app *Application) RegisterRoutes(e *echo.Echo) {
	e.GET("/", app.handler.Home)
	// Register more routes here
}

// Close handles closing the database connection
func (app *Application) Close() {
	database.Close() // Ensure to close the database connection
}
