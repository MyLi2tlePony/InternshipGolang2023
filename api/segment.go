package api

import "github.com/labstack/echo/v4"

type SegmentHandler interface {
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

func SetupSegmentRoutes(echo *echo.Echo, handler SegmentHandler) {
	echo.POST("/api/segment", handler.Create)
	echo.DELETE("/api/segment", handler.Delete)
}
