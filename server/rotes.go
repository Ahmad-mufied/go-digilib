package server

import (
	"github.com/Ahmad-mufied/go-digilib/server/handler"
	"github.com/Ahmad-mufied/go-digilib/server/middleware"
	"github.com/golang-jwt/jwt/v5"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.POST("/users/register", handler.Register)
	e.POST("/users/login", handler.Login)

	config := echoJWT.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.JWTCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	// Book Routes
	e.POST("/books", handler.CreateBook, echoJWT.WithConfig(config))
	e.GET("/books", handler.GetAllBooks, echoJWT.WithConfig(config))
	e.PUT("/books", handler.UpdateBookStock, echoJWT.WithConfig(config))
	e.GET("/books/:id", handler.GetBookDetails, echoJWT.WithConfig(config))

	// Deposit Routes
	e.POST("/deposits", handler.CreateDeposit, echoJWT.WithConfig(config))
	e.GET("/deposits/:id", handler.GetDepositById, echoJWT.WithConfig(config))

}
