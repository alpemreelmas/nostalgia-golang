package handlers

import "github.com/labstack/echo/v4"

// Handler struct for application handlers
type Handler struct{}

// Home is a simple handler for the root route
func (h *Handler) Home(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
