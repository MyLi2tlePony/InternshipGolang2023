package api

import "github.com/labstack/echo/v4"

type UserHandler interface {
	ChangeSegments(ctx echo.Context) error
	GetSegments(ctx echo.Context) error
}

func SetupUserRoutes(echo *echo.Echo, handler UserHandler) {
	echo.PATCH("/api/user/:id/segment", handler.ChangeSegments)
	echo.GET("/api/user/:id/segment", handler.GetSegments)
}
