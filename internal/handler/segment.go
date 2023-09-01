package handler

import (
	"fmt"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/domain"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type SegmentService interface {
	Create(segment *entity.Segment) error
	Delete(segment *entity.Segment) error
	CreateCSV(year, month int) (string, error)
}

type SegmentHandler struct {
	service SegmentService
}

func NewSegmentHandler(service SegmentService) *SegmentHandler {
	return &SegmentHandler{
		service: service,
	}
}

// Create
// @Param segment body domain.Segment true "Segment"
// @Success 200
// @Failure 400
// @Router /api/segment [post]
func (h *SegmentHandler) Create(ctx echo.Context) error {
	var body domain.Segment
	if err := ctx.Bind(&body); err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	entitySegment := &entity.Segment{
		Name: body.Name,
	}

	err := h.service.Create(entitySegment)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	return ctx.NoContent(http.StatusOK)
}

// Delete
// @Param segment body domain.Segment true "Segment"
// @Success 200
// @Failure 400
// @Router /api/segment [delete]
func (h *SegmentHandler) Delete(ctx echo.Context) error {
	var body domain.Segment
	if err := ctx.Bind(&body); err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	entitySegment := &entity.Segment{
		Name: body.Name,
	}

	err := h.service.Delete(entitySegment)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	return ctx.NoContent(http.StatusOK)
}

// CreateLinkCSV
// @Param segment body domain.CreateLinkCSVRequest true "period"
// @Success 200 {object} domain.CreateLinkCSVResponse
// @Failure 400
// @Router /api/segment/csv/ [post]
func (h *SegmentHandler) CreateLinkCSV(ctx echo.Context) error {
	var body domain.CreateLinkCSVRequest
	if err := ctx.Bind(&body); err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	fileName, err := h.service.CreateCSV(body.Year, body.Month)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	hostPort := ctx.Echo().ListenerAddr()
	return ctx.JSON(http.StatusOK, fmt.Sprintf("%s%s/%s", hostPort, ctx.Path(), fileName))
}

// GetCSV
// @Param name path string true "file name"
// @Success 200
// @Failure 400
// @Router /api/segment/csv/{name} [get]
func (h *SegmentHandler) GetCSV(ctx echo.Context) error {
	file := fmt.Sprintf("csv/%s", ctx.Param("name"))

	if _, err := os.Stat(file); err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	ctx.Response().Header().Set("content-disposition", "attachment; filename=\""+file+"\"")

	return ctx.File(file)
}
