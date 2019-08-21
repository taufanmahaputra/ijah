package service

import (
	"github.com/labstack/echo"
	"net/http"
)

type HTTPService struct {
}

func NewHTTPServer() HTTPService {
	return HTTPService{}
}

func (s HTTPService) RegisterHandler(e *echo.Echo) {
	e.GET("/", index)
}

func index(ctx echo.Context) error {
	return response(ctx, http.StatusOK, Response{
		"Ijah Inventory Management API",
		"OK",
	})
}

func response(ctx echo.Context, statusCode int, response interface{}) error {
	return ctx.JSON(statusCode, response)
}