package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) createUser(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

// upload and dwonload file
// create folder
// add file to folder
// remove file from folder
// delete folder

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
