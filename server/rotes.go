package server

import (
	"github.com/Ahmad-mufied/go-digilib/server/handler"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.POST("/users/register", handler.Register)
	e.POST("/users/login", handler.Login)

}
