package http

import (
	"net/http"
	"nuage/internal/repositories"

	"github.com/labstack/echo/v4"
)

func (s *Server) createUser(c echo.Context) error {

	var req struct {
		Email    string `json:"name"`
		FullName string `json:"fullname"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "bad request"})
	}

	db := &repositories.InMemoryUserRepository{}
	newUser, err := db.CreateUser(req.Email, req.Password, req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
	}

	resp := echo.Map{
		"status":  true,
		"message": "user created sucessfully",
		"data":    newUser,
	}

	return c.JSON(http.StatusOK, resp)
}

// upload and dowonload file
// create folder
// add file to folder
// remove file from folder
// delete folder

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
