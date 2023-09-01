package api

import "github.com/labstack/echo/v4"

type SegmentHandler interface {
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
	CreateLinkCSV(ctx echo.Context) error
	GetCSV(ctx echo.Context) error
}

func SetupSegmentRoutes(echo *echo.Echo, handler SegmentHandler) {
	echo.POST("/api/segment", handler.Create)
	echo.DELETE("/api/segment", handler.Delete)
	echo.POST("/api/segment/csv", handler.CreateLinkCSV)
	echo.GET("/api/segment/csv/:name", handler.GetCSV)
}
