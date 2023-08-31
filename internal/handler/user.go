package handler

import (
	"net/http"
	"strconv"

	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/domain"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	ChangeSegments(userID int, insert, delete []entity.Segment) error
	GetSegments(userID int) ([]entity.Segment, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// ChangeSegments
// @Param id path int true "User ID"
// @Param ChangeUserSegments body domain.ChangeUserSegments true "Changed User Segments"
// @Success 200
// @Failure 400
// @Router /api/user/{id}/segment [patch]
func (h *UserHandler) ChangeSegments(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	var body domain.ChangeUserSegments
	if err := ctx.Bind(&body); err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	var insertSegments, deleteSegments []entity.Segment
	for _, segment := range body.DeleteSegments {
		deleteSegments = append(deleteSegments, entity.Segment{Name: segment.Name})
	}

	for _, segment := range body.InsertSegments {
		insertSegments = append(insertSegments, entity.Segment{Name: segment.Name})
	}

	err = h.service.ChangeSegments(id, insertSegments, deleteSegments)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	return ctx.NoContent(http.StatusOK)
}

// GetSegments
// @Param id path int true "User ID"
// @Success 200 {object} []domain.Segment
// @Failure 400
// @Router /api/user/{id}/segment [get]
func (h *UserHandler) GetSegments(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	segments, err := h.service.GetSegments(id)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	domainSegments := make([]domain.Segment, 0, len(segments))
	for _, segment := range segments {
		domainSegments = append(domainSegments, domain.Segment{Name: segment.Name})
	}

	return ctx.JSON(http.StatusOK, domainSegments)
}
