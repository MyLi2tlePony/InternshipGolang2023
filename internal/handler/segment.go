package handler

import (
	"net/http"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/domain"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/labstack/echo/v4"
)

type SegmentService interface {
	Create(segment *entity.Segment) error
	Delete(segment *entity.Segment) error
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
